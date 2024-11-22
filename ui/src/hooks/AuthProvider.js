import { createContext, useContext, useEffect, useState } from "react"
import axios from "axios"

const AuthContext = createContext()

const AuthProvider = ({children}) => {
    const [user,setUser] = useState(null)
    const [token,setToken] = useState(localStorage.getItem("token_"+user?.username) || localStorage.getItem("token"))
    const apiHost = process.env.REACT_APP_API_URL

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
            axios.get(`${apiHost}/users`,{
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
    },[apiHost,token,user])

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