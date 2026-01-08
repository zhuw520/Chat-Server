let ws = null
let userId = Math.random().toString(36).substring(2) + Date.now()
let userId = localStorage.getItem('chat_user_id')
if (!userId) {
    userId = Math.random().toString(36).substring(2) + Date.now()
    localStorage.setItem('chat_user_id', userId)
}
function connect() {
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    const wsUrl = protocol + '//' + window.location.host + '/ws'
    ws = new WebSocket(wsUrl)
    
    ws.onopen = function() {
        document.getElementById('status').textContent = 'Connected'
        ws.send(JSON.stringify({type: 'join', userId: userId}))
    }
    
    ws.onmessage = function(event) {
        const data = JSON.parse(event.data)
        if (data.type === 'init') {
            data.messages.forEach(addMessage)
        } else if (data.type === 'chat_update') {
            addMessage(data.message)
        }
    }
    
    ws.onclose = function() {
        document.getElementById('status').textContent = 'Disconnected'
        setTimeout(connect, 3000)
    }
}

function addMessage(msg) {
    const div = document.createElement('div')
    div.className = 'message'
    
    if (msg.type === 'system' || msg.type === 'join') {
        div.textContent = msg.message
        div.style.color = '#666'
    } else if (msg.type === 'user') {
        div.innerHTML = 'User' + msg.userNumber + ': ' + msg.message
    }
    
    document.getElementById('messages').appendChild(div)
    document.getElementById('messages').scrollTop = document.getElementById('messages').scrollHeight
}

function sendMessage() {
    const input = document.getElementById('input')
    const text = input.value.trim()
    
    const now = new Date()
    const time = now.getHours().toString().padStart(2, '0') + ':' + now.getMinutes().toString().padStart(2, '0')
    
    ws.send(JSON.stringify({
        type: 'message',
        userId: userId,
        message: text,
        timestamp: time
    }))
    
    input.value = ''
}

document.getElementById('send').addEventListener('click', sendMessage)
document.getElementById('input').addEventListener('keypress', function(e) {
    if (e.key === 'Enter') sendMessage()
})

connect()