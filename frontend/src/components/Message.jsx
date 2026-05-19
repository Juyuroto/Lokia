import "../assets/css/message.css"

export default function Message({ role, content }) {
    return (
        <div className={`message ${role}`}>
            <span className="message-role">{role === "user" ? "Toi" : "Lokia"}</span>
            <p className="message-content">{content}</p>
        </div>
    )
}