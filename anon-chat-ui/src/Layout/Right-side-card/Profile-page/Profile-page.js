import img from '../../../assets/blank-dp.webp';
import { useAuth } from '../../../hooks/AuthProvider'
import './Profile-page.css'

const ProfilePage = () => {
    const userContext = useAuth()

    const handleLogout = () => {
        userContext.LogoutAction()
    }

    return (
        <div className="profile-page">
            <div className='profile'>
                <img src={img} alt=''/>
                <h3>{userContext.user.username}</h3>
            </div>
            <button className='logout-button' onClick={handleLogout}>Logout</button>
        </div>
    )
}

export default ProfilePage