import React, { useEffect, useState } from 'react'
import axios from 'axios';
import Rating from '@mui/material/Rating';
import {
    MdSentimentSatisfiedAlt,
    MdSentimentDissatisfied,
    MdSentimentVeryDissatisfied,
    MdSentimentNeutral,
    MdSentimentVerySatisfied,
    MdStarRate,
    MdOutlineSentimentVeryDissatisfied,
    MdSend,
} from 'react-icons/md'
import Box from '@mui/material/Box';
import { Button, MenuItem, Select, TextField, Tooltip, Typography } from '@mui/material';
import { toast } from 'react-toastify';
import './Review.css'
import CommentCard from '../Card/Comment Card/CommentCard';
import { customerReview } from '../../Assets/Images/Image';



const labels = {
    0: <MdOutlineSentimentVeryDissatisfied style={{ color: 'red' }} />,
    0.5: <MdOutlineSentimentVeryDissatisfied style={{ color: 'red' }} />,
    1: <MdSentimentVeryDissatisfied style={{ color: 'red' }} />,
    1.5: <MdSentimentVeryDissatisfied style={{ color: 'red' }} />,
    2: <MdSentimentDissatisfied style={{ color: 'orange' }} />,
    2.5: <MdSentimentDissatisfied style={{ color: 'orange' }} />,
    3: <MdSentimentNeutral style={{ color: 'gold' }} />,
    3.5: <MdSentimentNeutral style={{ color: 'gold' }} />,
    4: <MdSentimentSatisfiedAlt style={{ color: 'green' }} />,
    4.5: <MdSentimentSatisfiedAlt style={{ color: 'green' }} />,
    5: <MdSentimentVerySatisfied style={{ color: 'green' }} />,
};


function getLabelText(value) {
    return `${value} Star${value !== 1 ? 's' : ''}, ${labels[value]}`;
}
const ProductReview = ({ authToken, setProceed, setOpenAlert, id }) => {
    const [value, setValue] = useState(0);
    const [hover, setHover] = useState('');
    const [reviews, setReviews] = useState([]);
    const [comment, setComment] = useState('');
    const [title, setTitle] = useState('All');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);


    const fetchReviews = async () => {
        try {
            setLoading(true);
            setError(null);

            
            const response = await axios.get(`${process.env.REACT_APP_GET_REVIEW}/${id}/comments`);
            
            setReviews(response.data.comments || []);
        } catch (error) {
            console.error('Error fetching reviews:', error);
            setError(error.message);
            setReviews([]);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        if (id) {
            fetchReviews();
        }
    }, [id]);


    const handleSubmitReview = async (e) => {
        e.preventDefault();
        if (!comment && !value) {
            toast.error("Please Fill the all Fields", { theme: "colored", autoClose: 500 });
            return;
        }
        if (comment.length <= 4) {
            toast.error("Please add more than 4 characters", { theme: "colored", autoClose: 500 });
            return;
        }
        if (value <= 0) {
            toast.error("Please add rating", { theme: "colored", autoClose: 500 });
            return;
        }

        try {
            if (!setProceed) {
                setOpenAlert(true);
                return;
            }

            console.log('Submitting review:', { productId: id, content: comment, rating: value });
            const { data } = await axios.post(
                `${process.env.REACT_APP_ADD_IN_REVIEW}`,
                { productId: id, content: comment, rating: value },
                {
                    headers: {
                        'Authorization': `Bearer ${authToken}`,
                        'Content-Type': 'application/json'
                    }
                }
            );
            console.log('Review submission response:', data);
            toast.success(data.msg, { theme: "colored", autoClose: 500 });
            fetchReviews();
            setComment('');
            setValue(0);
        } catch (error) {
            console.error('Error submitting review:', error);
            toast.error(error.response?.data?.msg || "Error submitting review", { theme: "colored", autoClose: 600 });
            setComment('');
            setValue(0);
        }
    }
    return (
        <>
            <div className='form-container'>
                <form onSubmit={handleSubmitReview} className='form'>
                    <Box
                        sx={{
                            width: 300,
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'flex-start'
                        }}
                    >
                        <Rating
                            name="hover-feedback"
                            value={value}
                            precision={0.5}
                            getLabelText={getLabelText}
                            id="rating"
                            onChange={(event, newValue) => {
                                setValue(newValue);
                            }}
                            onChangeActive={(event, newHover) => {
                                setHover(newHover);
                            }}
                            emptyIcon={<MdStarRate style={{ opacity: 0.55 }} fontSize="inherit" />}
                        />
                        {value !== null && (
                            <Box className='expression-icon' sx={{ ml: 2 }}>{labels[hover !== -1 ? hover : value]}</Box>
                        )}
                    </Box>
                    <TextField
                        id="filled-textarea"
                        value={comment} onChange={(e) => {
                            setComment(e.target.value)
                        }}
                        label="Add Review"
                        placeholder="What did you like or dislike?"
                        multiline
                        className='comment'
                        variant="outlined"
                    />

                    <Tooltip title='Send Review'>
                        <Button className='form-btn' variant='contained' type='submit' endIcon={<MdSend />}>Send</Button>
                    </Tooltip>

                </form>
                <div className="form-img-box">
                    <img src={customerReview} loading='lazy' alt="Customer Review" className='review-img' />
                </div>
            </div>
            {reviews.length >= 1 ? <Box sx={{ display: 'flex', justifyContent: 'flex-end', gap: 1, width: "80vw" }}>
    
            </Box>
                :
                <Typography sx={{ textAlign: 'center' }}>No reviews have been submitted for this product yet. Be the first to add a review!</Typography>
            }
            <Box className='review-box' >
                {
                    reviews.map(review =>
                        <CommentCard userReview={review} key={review._id} authToken={authToken} setReviews={setReviews} reviews={reviews} fetchReviews={fetchReviews} />
                    )
                }
            </Box>

        </>
    )
}

export default ProductReview




