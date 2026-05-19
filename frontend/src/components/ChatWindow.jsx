import { useEffect, useState } from "react"
import { fetchConversationByID, sendMessage } from "../services/api.service"
import Message from "./Message"
import InputBar from "./InputBar"
import "../assets/css/chat.css"

export default function ChatWindow({ conversation }) {
  
    const [messages, setMessages] = useState([])
    const [loading, setLoading] = useState(false)
    
    useEffect(() => {
        if (!conversation) return
        fetchConversationByID(conversation.ID).then(data => setMessages(data))
    }, [conversation])

    const handleSend = async (prompt) => {
        if (!prompt.trim() || loading) return

        const userMessage = { Role: "user", Content: prompt }
        setMessages(prev => [...prev, userMessage])
        setLoading(true)

        const res = await sendMessage(conversation.ID, prompt)
        const iaMessage = { Role: "assistant", Content: res.response }
        setMessages(prev => [...prev, iaMessage])
        setLoading(false)
    }

    if (!conversation) return (
        <div className="chat-empty">
            <p>Sélectionne ou crée une conversation</p>
        </div>
    )

    return (
        <div className="chat-window">
            <div className="chat-header">
                <h2>{conversation.title}</h2>
            </div>
            <div className="chat-messages">
                {messages.map((msg, i) => (
                    <Message key={i} role={msg.Role} content={msg.Content} />
                ))}
                {loading && <Message role="assistant" content="..." />}
            </div>
            <InputBar onSend={handleSend} loading={loading} />
        </div>
    )
}