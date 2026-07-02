import './Friends-page.css';
import Friends from '../../Left-side-card/Friends/Friends';

const FriendsPage = ({onSwitch}) => {
    return(
        <div className="friends-page">
            <Friends onSwitch={onSwitch}/>
        </div>
    )
}

export default FriendsPage;