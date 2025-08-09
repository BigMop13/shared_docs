package httpserver

import (
	"html/template"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Shared Docs - Real-time Collaboration</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); min-height: 100vh; display: flex; flex-direction: column; }
        .header { background: rgba(255, 255, 255, 0.1); backdrop-filter: blur(10px); padding: 1rem 2rem; border-bottom: 1px solid rgba(255, 255, 255, 0.2); }
        .header h1 { color: white; font-size: 1.8rem; font-weight: 300; }
        .status { color: rgba(255, 255, 255, 0.8); font-size: 0.9rem; margin-top: 0.5rem; }
        .main-content { flex: 1; padding: 2rem; display: flex; flex-direction: column; }
        .editor-container { background: white; border-radius: 12px; box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1); overflow: hidden; flex: 1; display: flex; flex-direction: column; }
        .editor-header { background: #f8f9fa; padding: 1rem 1.5rem; border-bottom: 1px solid #e9ecef; display: flex; justify-content: space-between; align-items: center; }
        .editor-title { font-size: 1.2rem; color: #495057; font-weight: 500; }
        .connection-status { display: flex; align-items: center; gap: 0.5rem; }
        .status-dot { width: 8px; height: 8px; border-radius: 50%; background: #dc3545; transition: background 0.3s ease; }
        .status-dot.connected { background: #28a745; }
        .status-text { font-size: 0.9rem; color: #6c757d; }
        .text-editor { flex: 1; border: none; outline: none; padding: 2rem; font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace; font-size: 14px; line-height: 1.6; color: #212529; resize: none; background: #fafbfc; }
        .text-editor:focus { background: white; }
        .footer { background: rgba(255, 255, 255, 0.1); backdrop-filter: blur(10px); padding: 1rem 2rem; border-top: 1px solid rgba(255, 255, 255, 0.2); text-align: center; color: rgba(255, 255, 255, 0.8); font-size: 0.9rem; }
        .typing-indicator { color: #6c757d; font-style: italic; font-size: 0.9rem; }
        @media (max-width: 768px) { .main-content { padding: 1rem; } .text-editor { padding: 1rem; font-size: 16px; } }
    </style>
</head>
<body>
    <div class="header">
        <h1>📝 Shared Docs</h1>
        <div class="status">Real-time collaborative document editing</div>
    </div>
    <div class="main-content">
        <div class="editor-container">
            <div class="editor-header">
                <div class="editor-title">Document</div>
                <div class="connection-status">
                    <div class="status-dot" id="statusDot"></div>
                    <span class="status-text" id="statusText">Connecting...</span>
                </div>
            </div>
            <textarea class="text-editor" id="textEditor" placeholder="Start typing your document here...&#10;&#10;This is a real-time collaborative editor. Changes are automatically synchronized with all connected users." spellcheck="false"></textarea>
        </div>
    </div>
    <div class="footer">
        <div class="typing-indicator" id="typingIndicator"></div>
    </div>
    <script>
        let ws; let reconnectAttempts = 0; const maxReconnectAttempts = 5; let reconnectTimeout; let typingTimeout; let isTyping = false;
        const textEditor = document.getElementById('textEditor'); const statusDot = document.getElementById('statusDot'); const statusText = document.getElementById('statusText'); const typingIndicator = document.getElementById('typingIndicator');
        function connect() {
            const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'; const wsUrl = protocol + '//' + window.location.host + '/ws'; ws = new WebSocket(wsUrl);
            ws.onopen = function() { statusDot.classList.add('connected'); statusText.textContent = 'Connected'; reconnectAttempts = 0; };
            ws.onmessage = function(event) { try { const data = JSON.parse(event.data); if (data.content !== undefined) { if (textEditor.value !== data.content) { const cursorPos = textEditor.selectionStart; textEditor.value = data.content; textEditor.setSelectionRange(cursorPos, cursorPos); } } } catch (e) { console.error('Error parsing message:', e); } };
            ws.onclose = function() { statusDot.classList.remove('connected'); statusText.textContent = 'Disconnected'; if (reconnectAttempts < maxReconnectAttempts) { reconnectAttempts++; const delay = Math.min(1000 * Math.pow(2, reconnectAttempts), 10000); statusText.textContent = 'Reconnecting in ' + Math.round(delay/1000) + 's...'; reconnectTimeout = setTimeout(connect, delay); } else { statusText.textContent = 'Connection failed'; } };
            ws.onerror = function(error) { console.error('WebSocket error:', error); };
        }
        function debounce(func, wait) { let timeout; return function executedFunction(...args) { const later = () => { clearTimeout(timeout); func(...args); }; clearTimeout(timeout); timeout = setTimeout(later, wait); }; }
        const sendUpdate = debounce(function(content) { if (ws && ws.readyState === WebSocket.OPEN) { ws.send(JSON.stringify({ content: content })); } }, 100);
        textEditor.addEventListener('input', function() { if (!isTyping) { isTyping = true; typingIndicator.textContent = 'You are typing...'; } clearTimeout(typingTimeout); typingTimeout = setTimeout(() => { isTyping = false; typingIndicator.textContent = ''; }, 1000); sendUpdate(this.value); });
        textEditor.addEventListener('paste', function() { setTimeout(() => { sendUpdate(this.value); }, 10); });
        function autoResize() { textEditor.style.height = 'auto'; textEditor.style.height = textEditor.scrollHeight + 'px'; }
        textEditor.addEventListener('input', autoResize);
        connect();
        document.addEventListener('visibilitychange', function() { if (document.hidden) { if (reconnectTimeout) { clearTimeout(reconnectTimeout); } } else { if (ws && ws.readyState !== WebSocket.OPEN) { connect(); } } });
        window.addEventListener('beforeunload', function() { if (ws) { ws.close(); } });
    </script>
</body>
</html>`

	tmplParsed, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	_ = tmplParsed.Execute(w, nil)
}
