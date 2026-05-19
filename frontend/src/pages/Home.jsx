import { useState } from "react"
import Sidebar from "../components/Sidebar"
import ChatWindow from "../components/ChatWindow"

export default function Home() {
    const [activeConversation, setActiveConversation] = useState(null)

    return (
        <div className="home">
            <Sidebar
                onSelectConversation={setActiveConversation}
                activeConversationId={activeConversation?.ID}
            />
            <ChatWindow conversation={activeConversation} />
        </div>
    )
}