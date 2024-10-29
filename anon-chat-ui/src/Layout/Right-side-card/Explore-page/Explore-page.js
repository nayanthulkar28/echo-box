import './Explore-page.css'
import wand from '../../../assets/magic-wand.jpg'
import cancel from '../../../assets/cancel.png'
import { useState } from 'react';
import { useWebSocket } from '../../../hooks/Websocket';

const ExplorePage = () => {
    const [searching, setSearching] = useState(false)
    const ws = useWebSocket()

    const handleAddExplorer = () => {
        const req = {
            type: 2,
            to: null,
            message: null
        }
        ws.SendRequest(JSON.stringify(req))
        setSearching(true)
    }

    const handleRemoveExplorer = () => {
        const req = {
            type: 3,
            to: null,
            message: null
        }
        ws.SendRequest(JSON.stringify(req))
        setSearching(false)
    }

    return(
        <div className="explore-page">
            {!searching?
                <img src={wand} alt='' onClick={handleAddExplorer}/>:
                <img src={cancel} alt='' onClick={handleRemoveExplorer}/>
            }
        </div>
    );
}

export default ExplorePage;