import ProfileIcon from "./ProfileIcon";

function Story() {
   // let accountName = users[Math.floor(Math.random() * users.length)].username;

   // if (accountName.length > 10) {
    //    accountName = accountName.substring(0, 10) + "...";
   // }

    const story={
        display: 'flex',
        flexDirection: 'column',
         alignItems: 'center',
         margin: '1em 0.5em',
    }

    return (
        <div style={story}>
            <ProfileIcon iconSize="big" storyBorder={true} />
            <span className="accountName">'BLA'</span>
        </div>
    );
}

export default Story;