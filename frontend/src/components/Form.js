import React, { Component } from 'react';
 
class Form extends Component {
    //create refs
    authorRef = React.createRef();
    titleRef = React.createRef();
    contentRef = React.createRef();
    categoryRef = React.createRef();
 
 
    createPost = (e) => {
        e.preventDefault();
 
        const post = {
            author: this.authorRef.current.value,
            title: this.titleRef.current.value,
            body: this.contentRef.current.value,
            category: this.categoryRef.current.value
        }
 
        this.props.createPost(post);
 
    }
 
 
    render() { 
        return ( 
            <form onSubmit={this.createPost} className="col-md-10">
                <legend className="text-center">Create New Post</legend>
 
                <div className="form-group">
                    <label>Title for the Post:</label>
                    <input type="text" ref={this.titleRef} className="form-control" placeholder="Title.." />
                </div>
 
                <div className="form-group">
                    <label>Author:</label>
                    <input type="text" ref={this.authorRef} className="form-control" placeholder="Tag your name.." />
                </div>

                <div className="col">
                <input
                value={location}
                onChange={(e) => setLocation(e.target.value)}
                className="form-control"
                type="text"
                placeholder="location"
                />
                </div>

                <div className="col">
                <select
                value={priceRange}
                onChange={(e) => setPriceRange(e.target.value)}
                className="custom-select my-1 mr-sm-2"
                >
                <option disabled>Price Range</option>
                <option value="1">$</option>
                <option value="2">$$</option>
                <option value="3">$$$</option>
                <option value="4">$$$$</option>
                <option value="5">$$$$$</option>
                </select>
                </div>
                
                <div className="form-group">
                    <label>Content:</label>
                    <textarea className="form-control" rows="7"cols="25" ref={this.contentRef} placeholder="Here write your content.."></textarea>
                </div>
 
                <div className="form-group">
                    <label>Category</label>
                <select ref={this.categoryRef} className="form-control" id="category">
                    <option value="adventure">Adventure</option>
                    <option value="airport">Airport</option>
                    <option value="bar">Bar</option>
                    <option value="cafe">Cafe</option>
                    <option value="club">Club</option>
                    <option value="Downtown">Downtown</option>
                    <option value="hotel">Hotel</option>
                    <option value="lake">Lake</option>
                    <option value="library">Library</option>
                    <option value="mall">Mall</option>
                    <option value="midtown">Midtown</option>
                    <option value="motel">Motel</option>
                    <option value="museum">Museum</option>
                    <option value="nature">Nature</option>
                    <option value="park">Park</option>
                    <option value="planetarium">Planetarium</option>
                    <option value="restaurant">Restaurant</option> 
                    <option value="santafevent">Santa Fe event</option>
                    <option value="sport">Sports</option>
                    <option value="theatre">Theatre</option>
                    <option value="ufevent">UF event</option>
                    <option value="zoo">Zoo</option>
                </select>
                </div>
                
                <div className="form-group">
                    <label>Rating</label>
                <select ref={this.categoryRef} className="form-control" id="rate">
                    <option value="1">1</option>
                    <option value="2">2</option>
                    <option value="3">3</option>
                    <option value="4">4</option>
                    <option value="5">5</option>
                </select>
                </div>

                <input type="file" id="image-input" accept="image/jpeg, image/png, image/jpg"></input>
                <div id="display-image"></div>
                <button type="submit" className="btn btn-primary">Create</button>
            </form>
         );
    }
}
 
export default Form;
