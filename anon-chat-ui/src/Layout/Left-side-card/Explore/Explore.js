import './Explore.css'
import { FaMagic } from "react-icons/fa";

const Explore = ({onSwitch}) => {
    return(
        <div className="explore" onClick={()=>onSwitch("explore","")}>
            <h3>Explore</h3>
            <FaMagic/>
        </div>
    );
}

export default Explore;