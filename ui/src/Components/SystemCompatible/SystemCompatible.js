import './SystemCompatible.css'
import errorImg from '../../assets/error.png'

const SystemCompatible = ({children}) => {
    const width = window.innerWidth
    if(width > 1400) return children
    return (
        <div className='container'>
            <div className='holder'>
                <img src={errorImg}/>
                <p>Sorry, currently app is not supported for smaller devices. <br />
                Please visit the site in desktop.</p>
            </div>
        </div>
    )
}

export default SystemCompatible