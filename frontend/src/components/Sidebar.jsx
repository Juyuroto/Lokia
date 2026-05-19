import { useEffect, useState } from "react"
import { fetchConversations, createConversation, deleteConversation } from "../services/api.service"
import "../assets/css/sidebar.css"

export default function Sidebar({ onSelectConversation, activeConversationId }) {
    const [conversations, setConversations] = useState([])
    
    useEffect(() => {
        fetchConversations().then(data => setConversations(data))
    }, [])

    const handleCreate = async () => {
        const conv = await createConversation("Nouvelle conversation")
        setConversations([conv, ...conversations])
        onSelectConversation(conv)
    }

    const handleDelete = async (e, id) => {
        e.stopPropagation()
        await deleteConversation(id)
        setConversations(conversations.filter(c => c.ID !== id))
    }

    return (
        <aside className="sidebar">
            <div className="sidebar-header">
                <h1 className="sidebar-title">Lokia</h1>
                <button className="sidebar-new-btn" onClick={handleCreate}>+</button>
            </div>
            <ul className="sidebar-list">
                {conversations.map(conv => (
                    <li
                        key={conv.ID}
                        className={`sidebar-item ${activeConversationId === conv.ID ? "active" : ""}`}
                        onClick={() => onSelectConversation(conv)}
                    >
                        <span>{conv.title}</span>
                        <button className="sidebar-delete-btn" onClick={(e) => handleDelete(e, conv.ID)}>×</button>
                    </li>
                ))}
            </ul>
        </aside>
    )
}