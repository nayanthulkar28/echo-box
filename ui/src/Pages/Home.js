import { useState } from 'react';
import LeftCard from '../Layout/Left-side-card/Left-side-card';
import RightCard from '../Layout/Right-side-card/Right-side-card';
import TitleBar from '../Layout/TitleBar/TitleBar';
import './Home.css'
import { useWebSocket } from '../hooks/Websocket';

const Home = () => {
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
            <TitleBar onSwitch={handleComponentSwitch}/>
            <div className='home-body'>
                <LeftCard onSwitch={handleComponentSwitch}/>
                <RightCard component={rightComponent}/>
            </div>
        </div>
    );
}

export default Home;