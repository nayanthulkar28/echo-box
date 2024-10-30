import Explore from './Explore/Explore';
import './Left-side-card.css'
import Friends from './Friends/Friends';

const LeftCard = ({onSwitch}) => {
    return(
        <div className="left-card">
           <Explore onSwitch={onSwitch}/>
           <Friends onSwitch={onSwitch}/>
        </div>
    );
}

export default LeftCard;