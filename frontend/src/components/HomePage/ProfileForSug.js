import "../../style/profileSug.css";
import ProfileIcon from "../ProfileComponent/ProfileIcon";
import {Button, NavLink} from "react-bootstrap";
import {useHistory} from "react-router-dom";

function ProfileForSug(props) {
    const {user, username, caption,  urlText, iconSize,captionSize, storyBorder,hideAccountName,image, firstName, lastName} = props;
    const history = useHistory()

    function redirect(){
        history.push({
            pathname: '/profile/'+username,
        })
    }
    return (
        <div className="profile">
            <ProfileIcon iconSize={iconSize} storyBorder={storyBorder} image={image} />
            {(username || caption) && !hideAccountName && (
                <div className="textContainer">
                        <span className="accountName">{firstName} {lastName} @{username}</span>
                    <Button variant="link" onClick={redirect}>See profile</Button>
                    <span className={`caption ${captionSize}`}>{caption}</span>
                </div>
            )}
        </div>
    );
}

export default ProfileForSug;