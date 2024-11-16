import img from '../../../assets/blank-dp.webp';
import { useAuth } from '../../../hooks/AuthProvider'
import './Profile-page.css'

const ProfilePage = () => {
    const userContext = useAuth()

    return (
        <div className="profile-page">
            <div className='profile'>
                <img src={img} alt=''/>
                <h3>{userContext.user.username}</h3>
            </div>
        </div>
    )
}

export default ProfilePage