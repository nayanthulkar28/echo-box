import { useAuth } from '../../../hooks/AuthProvider'
import './NavBar.css'

const NavBar = ({onSwitch}) => {
    const userContext = useAuth()

    const handleLogOut = () => {
        userContext.LogoutAction()
    }

    return(
        <ul className="nav-bar">
            <li>Home</li>
            <li onClick={() => onSwitch("profile", "")}>Profile</li>
            <li>Setting</li>
            <li onClick={handleLogOut}>Sign Out</li>
        </ul>
    )
}

export default NavBar