import { useState } from 'react';
import ChatsPage from './Chats-page/Chats-page';
import ExplorePage from './Explore-page/Explore-page';
import ProfilePage from './Profile-page/Profile-page';
import './Right-side-card.css'
import { useWebSocket } from '../../hooks/Websocket';
import FriendsPage from './Friends-page/Friends-page';

const RightCard = ({component, onSwitch}) => {
    const [messages, setMessages] = useState({})
    const ws = useWebSocket()

    const appendMessage = (friend, message) => {
        setMessages((prev) => {
            if(prev[friend]) {
                return {
                    ...prev,
                    [friend]: [message, ...prev[friend]]
                }
            }
            return {
                ...prev,
                [friend]: [message]
            }
        })
    }

    ws.RegisterCallback("receive-message", appendMessage)

    return(
        <div className="right-card">
            {component.component.name==="explore" && <ExplorePage/>}
            {component.component.name==="chat" && 
                <ChatsPage 
                    friend={component.component.friendName}
                    isFriend={component.component.isFriend}
                    messages={messages[component.component.friendName]?messages[component.component.friendName]:[]}
                    appendMessage={appendMessage}/>}
            {component.component.name==="profile" && <ProfilePage/>}
            {component.component.name==="friends" && <FriendsPage onSwitch={onSwitch}/>}
        </div>
    );
}

export default RightCard;