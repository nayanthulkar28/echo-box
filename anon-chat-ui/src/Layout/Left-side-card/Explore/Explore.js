import './Explore.css'

const Explore = ({onSwitch}) => {
    return(
        <div className="explore" onClick={()=>onSwitch("explore","")}>
            Explore
        </div>
    );
}

export default Explore;