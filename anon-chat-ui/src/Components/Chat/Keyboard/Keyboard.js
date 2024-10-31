import { useState } from 'react';
import './Keyboard.css'
import { IoSend } from "react-icons/io5";
import { useWebSocket } from '../../../hooks/Websocket';

const Keyboard = ({friend, appendMessage}) => {
    const [message, setMessage] = useState('')
    const ws = useWebSocket()

    const handleSend = (e) => {
        e.preventDefault()
        if(message==="") return
        const req = {
            to: {
                user: {
                    username: friend
                }
            },
            type: 1,
            message: message
        }
        ws.SendRequest(JSON.stringify(req))
        appendMessage(friend, {
            value: e.target.message.value,
            send: 1
        })
        setMessage('')
    }

    return(
        <div className="keyboard">
            <form onSubmit={handleSend}>
                <textarea name='message' value={message} placeholder='Write a message ...' onChange={(e)=>setMessage(e.target.value)}/>
                <button type='submit'>
                    <IoSend/>
                </button>
            </form>
        </div>
    );
}

export default Keyboard