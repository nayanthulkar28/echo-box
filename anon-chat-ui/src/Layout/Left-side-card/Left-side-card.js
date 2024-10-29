import Explore from './Explore/Explore';
import Title from './Title/Title';
import './Left-side-card.css'
import Friends from './Friends/Friends';
import Menu from './Menu/Menu';

const LeftCard = ({onSwitch}) => {
    return(
        <div className="left-card">
           <Title/>
           <Explore onSwitch={onSwitch}/>
           <Friends onSwitch={onSwitch}/> 
           <Menu onSwitch={onSwitch}/>
        </div>
    );
}

export default LeftCard;