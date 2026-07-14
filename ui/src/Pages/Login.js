import { useEffect, useState } from 'react';
import Title from '../Layout/TitleBar/Title/Title'
import './Login.css'
import axios from "axios"
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../hooks/AuthProvider';

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate=useNavigate()
    const userContext = useAuth()
    const apiHost = process.env.REACT_APP_API_URL
    const [showPopUp, setShowPopUp] = useState(false)
    const [popUpMessage, setPopUpMessage] = useState("")

    useEffect(() => {
        if(userContext.user != null) {
            navigate("/", {replace: true})
        }
    })

    const handleSignUp = () => {
        axios.post(`${apiHost}/sign-up`,
            {
                username: username,
                password: password
            }
        )
        .then((res)=>{
            setUsername("")
            setPassword("")
            setShowPopUp(true)
            setPopUpMessage(res.data?.status)
            setTimeout(()=>setShowPopUp(false),3000)
        })
        .catch((error)=>{
            setShowPopUp(true)
            setPopUpMessage(error?.response.data.error)
            setTimeout(()=>setShowPopUp(false),3000)
            console.log(error.response)
        })
    }

    const handleLogin = () => {
        axios.post(`${apiHost}/login`,
            {
                username: username,
                password: password
            }
        ).then((res)=>{
            userContext.LoginAction(res.data)
            setUsername("")
            setPassword("")
            navigate("/")
        })
        .catch((error)=>{
            console.log(error)
        })
    }

    return(
        <div className='login-container'>
            {showPopUp && <div className="pop-up">{popUpMessage}</div>}
            <div className="login-form">
                <div className='login-title'>
                    <Title class_name="title-login"/>
                </div>
                <input type='text' name='username' placeholder='username' value={username} onChange={(e)=>setUsername(e.target.value)}/>
                <input type='text' name='password' placeholder='password' value={password} onChange={(e)=>setPassword(e.target.value)}/>
                <div className='button-container'>
                    <button type='submit' onClick={handleSignUp}>Sign Up</button>
                    <button type='submit' onClick={handleLogin}>Login</button>
                </div>
            </div>
        </div>
    );
}

export default Login