import { useState } from "react"

export default function InputBar({ onSend, loading }) {
    const [input, setInput] = useState("")

    const handleSend = () => {
        onSend(input)
        setInput("")
    }

    const handleKey = (e) => {
        if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault()
            handleSend()
        }
    }

    return (
        <div className="input-bar">
            <textarea
                className="input-textarea"
                value={input}
                onChange={e => setInput(e.target.value)}
                onKeyDown={handleKey}
                placeholder="Envoie un message..."
                rows={1}
                disabled={loading}
            />
            <button
                className="input-send-btn"
                onClick={handleSend}
                disabled={loading || !input.trim()}
            >
                Envoyer
            </button>
        </div>
    )
}