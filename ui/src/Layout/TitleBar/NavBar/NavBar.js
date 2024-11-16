import { useState } from 'react'
import { useAuth } from '../../../hooks/AuthProvider'
import './NavBar.css'

const NavBar = ({onSwitch}) => {
    const [focus, setFocus] = useState('chat')
    const userContext = useAuth()

    const handleLogOut = () => {
        userContext.LogoutAction()
    }

    const handleCLick = (focus) => {
        setFocus(focus)
    }

    return(
        <ul className="nav-bar">
            <li className={(focus==="chat"?"focus":"not-focus") + " nav-bar-items"} onClick={() => {onSwitch("explore", "");handleCLick("chat")}}>Chat</li>
            <li className={(focus==="profile"?"focus":"not-focus") + " nav-bar-items"} onClick={() => {onSwitch("profile", "");handleCLick("profile")}}>Profile</li>
            <li className={(focus==="settings"?"focus":"not-focus") + " nav-bar-items"}>Settings</li>
            <li className={"not-focus nav-bar-items"}onClick={handleLogOut}>Sign Out</li>
        </ul>
    )
}

export default NavBar