import './Friends.css'
import img from '../../../assets/blank-dp.webp'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { useWebSocket } from '../../../hooks/Websocket'
import { useAuth } from '../../../hooks/AuthProvider'

const Friends = ({onSwitch}) => {
    const ws = useWebSocket()
    const userContext = useAuth()
    const token = userContext.token
    const [friends, setFriends] = useState([])
    const friendList=[]
    friends.forEach(friend=>{
        friendList.push(
        <div className='friend-card-container' key={friend.username}>
            <div className="friend-card" onClick={()=>onSwitch("chat",friend.username)}>
                <img src={img} alt=''/>
                {friend.username}
            </div>
            <span className='divider'></span>
        </div>)
    })

    useEffect(() => {
        axios.get("http://localhost:8080/api/v1/users/friends",{
            headers: {
                "Authorization": "Bearer " + token
            }}
        )
        .then((res) => {
            setFriends(res.data)
        })
        .catch((error)=>{
            console.log(error)
        })
    },[token])

    ws.RegisterCallback("add-friend", (friend) => {
        setFriends((prev) => [friend, ...prev])
    })

    return(
        <div className="friends-container">
            <h3>Friends</h3>
            <div className="friend-card-list">{friendList}</div>
        </div>
    );
}

export default Friends;