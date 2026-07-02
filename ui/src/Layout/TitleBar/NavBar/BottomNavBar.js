import './BottomNavBar.css';
import { FaSearch, FaUserFriends } from "react-icons/fa";
import { IoIosSettings } from "react-icons/io";

const BottomNavBar = ({onSwitch}) => {
    return(
        <ul className="bottom-nav-bar">
            <li className="bottom-nav-bar-item" onClick={()=>onSwitch("explore","")}>
                <FaSearch />
            </li>
            <li className="bottom-nav-bar-item" onClick={()=>onSwitch("friends","")}>
                <FaUserFriends />
            </li>
            <li className="bottom-nav-bar-item" onClick={()=>onSwitch("profile","")}>
                <IoIosSettings />
            </li>
        </ul>
    )
}

export default BottomNavBar;