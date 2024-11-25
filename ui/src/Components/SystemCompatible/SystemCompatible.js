import './SystemCompatible.css'
import errorImg from '../../assets/error.png'
import Title from '../../Layout/TitleBar/Title/Title'

const SystemCompatible = ({children}) => {
    const width = window.innerWidth
    if(width > 1400) return children
    return (
        <div className='container'>
            <div className='holder'>
                <Title class_name="title-login"/>
                <img src={errorImg} alt=''/>
                <p>Sorry, currently app is not supported for smaller devices. <br />
                Please visit the site in desktop.</p>
            </div>
        </div>
    )
}

export default SystemCompatible