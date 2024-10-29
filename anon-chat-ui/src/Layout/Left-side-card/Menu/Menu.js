import './Menu.css'
import { CgProfile } from "react-icons/cg";
import { IoSettingsOutline } from "react-icons/io5";

const Menu = ({onSwitch}) => {
    return(
        <div className="Menu">
            <div className='Profile' onClick={() => onSwitch("profile","")}>
                <CgProfile className='Menu-items'/>
            </div>
            <div className='Settings'>
                <IoSettingsOutline className='Menu-items'/>
            </div>
        </div>
    );
}

export default Menu;