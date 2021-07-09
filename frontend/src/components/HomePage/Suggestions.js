import "../../style/suggestions.css";
import ProfileForSug from "./ProfileForSug";
import React, {useEffect, useState} from "react";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import followersService from "../../services/followers.service";
import {useSelector} from "react-redux";

function Suggestions() {
    const[recommendations,setRecommendations]=useState([])
    const store = useSelector(state => state);

    useEffect(() => {
        getRecommendations()
    }, []);

    async function getRecommendations() {
        const response = await followersService.getRecommendations(
            {
                id:store.user.id,
                jwt: store.user.jwt
            })

        if (response.status === 200) {
           console.log("JEEEJ")
           console.log(response.data)
        } else {
            toastService.show("error", "GRESKAA!");
        }
    }
    return (
        <div className="suggestions">
            <div className="titleContainer">
                <div className="title">Suggestions For You</div>
                <a href="/">See All</a>
            </div>

            <ProfileForSug username="maki" firstName="Marko" lastName="Markovic"  caption="Health" urlText="Follow" iconSize="medium" captionSize="small"  storyBorder={true} />
            <ProfileForSug username="joca" firstName="Jovan" lastName="Petrovic"  caption="Sport" urlText="Follow"  iconSize="medium" captionSize="small" />
            <ProfileForSug username="boki123" firstName="Bojana" lastName="Zoric" caption="Follows you" urlText="Follow"  iconSize="medium" captionSize="small" />
            <ProfileForSug username="majanokti123" firstName="Maja" lastName="Lazic" caption="Baby"  urlText="Follow" iconSize="medium"  captionSize="small" storyBorder={true} />
            <ProfileForSug username="lola" firstName="Jovana" lastName="Jokic" caption="Sport" urlText="Follow" iconSize="medium"  captionSize="small" />
        </div>
    );
}

export default Suggestions;