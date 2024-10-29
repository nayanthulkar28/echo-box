import { useState } from 'react';
import LeftCard from '../Layout/Left-side-card/Left-side-card';
import RightCard from '../Layout/Right-side-card/Right-side-card';
import './Home.css'
import { useWebSocket } from '../hooks/Websocket';

const Home = () => {
    console.log('home')
    const ws = useWebSocket()
    const [rightComponent, setRightComponent] = useState(
        {
            component: {
                name: "explore",
                friendName: "",
            }
        }
    );

    const handleComponentSwitch = (name, friendName) => {
        setRightComponent({
            component: {
                name: name,
                friendName: friendName,
                isFriend: true
            }
        });
    }

    ws.RegisterCallback("explore-chat", (friendName, isFriend) => {
        setRightComponent({
            component: {
                name: "chat",
                friendName:  friendName,
                isFriend: isFriend
            }
        })
    })

    return(
        <div className='home'>
            <LeftCard onSwitch={handleComponentSwitch}/>
            <RightCard component={rightComponent}/>
        </div>
    );
}

export default Home;