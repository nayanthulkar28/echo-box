import './Chat.css'
import img from '../../assets/blank-dp.webp'
import Keyboard from './Keyboard/Keyboard';
import { useState } from 'react';
import { useWebSocket } from '../../hooks/Websocket';

const Chat = ({friend, isFriend, messages, appendMessage}) => {
    const ws = useWebSocket()
    const messageList = []
    const [status, setStatus] = useState("Add Friend")

    messages.forEach((message,idx)=>{    
        messageList.push(
            <div key={idx} className="chat-message-container">
                <div className={"chat-message " + ((message.send===1)?"float-right":"float-left")}>{message.value}</div>
            </div>
        )
    })

    const handleButtonStatus = () => {
        const req = {
            to: {
                user: {
                    username: friend
                }
            }
        }
        if(status==="Accept Invite") {
            req.type=5
            ws.SendRequest(JSON.stringify(req))
            setStatus("Accepted")
        } else {
            req.type=4
            ws.SendRequest(JSON.stringify(req))
            setStatus("Request Sent")
        }
    }

    ws.RegisterCallback("receive-friend-request", () => {
        setStatus("Accept Invite")
    })

    return(
        <div className="chat-container">
            <div className="chat-title">
                <img src={img} alt=''/>
                <h5>{friend}</h5>
                {!isFriend && <button onClick={handleButtonStatus}>{status}</button>}
            </div>
            <div className="chat-body">
                {messageList}
            </div>
            <Keyboard friend={friend}
                appendMessage={appendMessage}/>
        </div>
    );
}

export default Chat;