import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import moment from 'moment';

import Navigation from "../HomePage/Navigation";
import { groupBy } from '../../util';
import monitoringService from '../../services/monitoring.service';

import './../../style/UserActivity.css'

const UserActivity = () => {
    const [activity, setActivity] = useState({})

    useEffect(() => {
        /* const tempActivity = [
            { id: "1", timestamp: moment("30-08-2021 15:30", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
            { id: "2", timestamp: moment("30-08-2021 15:36", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
            { id: "3", timestamp: moment("31-08-2021 08:14", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
            { id: "4", timestamp: moment("01-09-2021 09:15", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
            { id: "5", timestamp: moment("31-08-2021 12:20", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
            { id: "6", timestamp: moment("30-08-2021 12:20", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
            { id: "7", timestamp: moment("30-08-2021 13:30", "DD-MM-YYYY HH:mm"), message: "Failed login attempt." },
        ] */
        (async function () {
            let tempActivity = []
            const response = await monitoringService.getUsersRecentActivity({ jwt: "", userId: "agentagentic@gmail.com" })
            if (response.status == 200) {
                tempActivity = [...response.data]
            }
            console.log(tempActivity)

            const activityGroup = tempActivity.map(item => {
                return {
                    ...item,
                    date: moment(item.timestamp).format("D MMM YYYY"),
                    time: moment(item.timestamp).format("HH:mm")
                }
            })
            const grouped = groupBy(activityGroup, 'date');

            const keyOrder = Object.keys(grouped).sort(function (first, second) {
                if (moment(first.date).isAfter(moment(second.date))) return 1;
                return -1;
            })
            const result = {};
            keyOrder.forEach(key => {
                if (grouped[key].length > 1) {
                    let sorted = [...grouped[key].sort(function (first, second) {
                        const firstParts = first.time.split(":")
                        const secondParts = second.time.split(":")

                        if (Number(firstParts[0]) > Number(secondParts[0])) return -1;
                        else if (Number(firstParts[0]) < Number(secondParts[0])) return 1
                        else {
                            if (Number(firstParts[1]) > Number(secondParts[1])) return -1;
                            else if (Number(firstParts[1]) < Number(secondParts[1])) return 1
                        }

                        return 0;
                    })]
                    result[key] = [...sorted]
                } else {
                    result[key] = [...grouped[key]];
                }
            })

            setActivity(result)
        })()
    }, [])

    const activityByDate = (date) => {
        return <div className="date">
            <div className="title">{date}</div>
            <div className="items">{activity[date].map(item => activityItem(item))}</div>
        </div>
    }

    const activityItem = (item) => {
        return <div className="item">
            <div className="time">{item.time}</div>
            <div className="message">{item.message}</div>
        </div>
    }

    return (
        <div>
            <Navigation />
            <div className="UserActivity__Wrapper">
                <div className="title"> Recent Activity </div>
                <div className="listing">
                    {Object.keys(activity).map(date => activityByDate(date))}
                </div>
            </div>
        </div>
    )
}

export default UserActivity;