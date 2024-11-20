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

    useEffect(() => {
        if(userContext.user != null) {
            navigate("/", {replace: true})
        }
    })

    const handleSignUp = () => {
        axios.post("http://localhost:8090/echo-box-be/api/v1/sign-up",
            {
                username: username,
                password: password
            }
        ).catch((error)=>{
            console.log(error)
        })
        setUsername("")
        setPassword("")
    }

    const handleLogin = () => {
        axios.post("http://localhost:8090/echo-box-be/api/v1/login",
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