import React, {useState, useEffect} from 'react'
import axios from 'axios'
import Clamp from 'react-multiline-clamp';
import {Nav} from 'react-bootstrap'
import { formatDate } from "../utils";
import img1 from '../../assets/asset-1.png'

export default function HomePage() {
    const [posts, setPosts] = useState([])
    const [profiles, setProfiles] = useState([])
    const [postUsernames, setPostUsernames] = useState([])

    

    useEffect(() => {
        axios.get(axios.defaults.baseURL + "posts").then(async (res) => {
            setPosts(res.data)
            const timer = ms => new Promise(res => setTimeout(res, ms))
            
            for (var i = 0; i < res.data.length; i++) {
                var authorId = res.data[i].authorid
                console.log(authorId);

                axios.get(axios.defaults.baseURL + `profiles/${authorId}`).then(r => {
                    setPostUsernames(postUsernames => [...postUsernames, r.data.name])
                })

                await timer(100)
            }
        })
    }, [])

    console.log(postUsernames);

    return (
        <div>
            <div style={{width: "65%", margin: "auto", marginTop: 72}}>
                <h1 style={{fontSize:80 , textAlign: "center", marginBottom: 64}}>GatorX</h1>
                <div>
                    <h2 style={{textAlign: "center"}}>Latest posts</h2>
                    <hr />
                    <div>
                        {
                            posts.map((item, index) => 
                                <Nav.Link href={`/posts/${item.id}`} style={{marginTop: 48, color: "black"}}>
                                    <h4 style={{textAlign: "center"}}>{item.title}</h4>
                                    <p style={{textAlign: "center"}}>
                                        <small>{formatDate(item.createdat)} {String.fromCharCode(183)} {postUsernames[index]}</small>
                                        </p>
                                    <Clamp withTooltip lines={6}>
                                            <p style={{marginTop: 32, fontSize: 14}}>
                                                {item.content}
                                            </p>
                                        </Clamp>
                                </Nav.Link>
                            )
                        }
                    </div>
                </div>
            </div>
        </div>
        <div className="container hero">
                    <div className="row align-items-center text-center text-md-left">
                        <div className="col-lg-4">
                            <h1 className="mb-3 display-3">
                                Share your experience with the world!
                    </h1>
                            <p>
                                Join us! Login or Register. Write your review and share !!
                    </p>
                        </div>
                        <div className="col-lg-8">
                            <img src={img1} className="img-fluid" alt="img" />
                        </div>
                    </div>


                </div>
    )
}
