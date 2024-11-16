import { createContext, useContext, useEffect, useState } from "react"
import axios from "axios"

const AuthContext = createContext()

const AuthProvider = ({children}) => {
    const [user,setUser] = useState(null)
    const [token,setToken] = useState(localStorage.getItem("token_"+user?.username) || localStorage.getItem("token"))

    const LoginAction = (data) => {
        setUser(data.user)
        if(data?.token) {
            setToken(data.token)
            localStorage.setItem("token",data.token)
            localStorage.setItem("token_"+data.user.username,data.token)
        }
    }
    const LogoutAction = () => {
        localStorage.removeItem("token")
        localStorage.removeItem("token_"+user.username)
        setUser(null)
        setToken(null)
    }

    useEffect(()=>{
        if(token!=null && user==null) {
            axios.get("http://localhost:8090/echo-box/api/v1/users",{
                headers: {
                    "Authorization": "Bearer " + token
                }
            })
            .then((res) => {
                LoginAction(res.data)
            })
            .catch((error) => {
                console.log(error)
            })
        }
    },[token,user])

    return(
        <AuthContext.Provider value={{user,token,LoginAction,LogoutAction}}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthProvider;

export const useAuth = ()=> {
    return useContext(AuthContext)
}