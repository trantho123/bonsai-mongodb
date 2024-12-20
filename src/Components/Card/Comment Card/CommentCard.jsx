import { Avatar, Box, Button, Grid, Rating, SpeedDial, SpeedDialAction, SpeedDialIcon, TextField } from '@mui/material'
import axios from 'axios';
import { useEffect, useState } from 'react';
import { AiFillEdit, AiFillDelete, AiOutlineSend } from 'react-icons/ai'
import { GiCancel } from 'react-icons/gi'
import { toast } from 'react-toastify';

const CommentCard = ({ userReview, setReviews, reviews, fetchReviews }) => {
    const formatDate = (dateString) => {
        if (!dateString) return 'N/A';
        const date = new Date(dateString);
        if (isNaN(date.getTime())) return 'Invalid Date';
        return date.toLocaleDateString('en-us', { 
            weekday: "long", 
            year: "numeric", 
            month: "short", 
            day: "numeric" 
        });
    }

    const formatTime = (dateString) => {
        if (!dateString) return 'N/A';
        const date = new Date(dateString);
        if (isNaN(date.getTime())) return 'Invalid Time';
        return date.toLocaleTimeString('en-US');
    }

    const [authUser, setAuthUser] = useState()
    const [editComment, setEditComment] = useState(userReview.content)
    const [edit, setEdit] = useState(false)
    const [isAdmin, setIsAdmin] = useState(false)
    const [value, setValue] = useState(userReview.rating);
    let authToken = localStorage.getItem('Authorization')

    useEffect(() => {
        authToken && getUser()
    }, [])

    const getUser = async () => {
        try {
            const { data } = await axios.get(`${process.env.REACT_APP_GET_USER_DETAILS}`, {
                headers: {
                    'Authorization': `Bearer ${authToken}`
                }
            })
            setAuthUser(data.data.email);
            setIsAdmin(data.data.role === "admin");
        } catch (error) {
            console.error('Error fetching user:', error);
        }
    }

    const handleDeleteComment = async () => {
        try {
            const { data } = await axios.delete(`${process.env.REACT_APP_DELETE_REVIEW}/${userReview._id}`, {
                headers: {
                    'Authorization': `Bearer ${authToken}`
                }
            })
            toast.success(data.msg, { autoClose: 500, theme: 'colored' })
            setReviews(reviews.filter(r => r._id !== userReview._id))
        } catch (error) {
            toast.success(error, { autoClose: 500, theme: 'colored' })
        }

    }
    const deleteCommentByAdmin = async () => {
        if (isAdmin) {
            try {
                const { data } = await axios.delete(`${process.env.REACT_APP_ADMIN_DELETE_REVIEW}/${userReview._id}`,
                    {
                        headers: {
                            'Authorization': `Bearer ${authToken}`
                        }
                    })
                toast.success(data.msg, { autoClose: 500, theme: 'colored' })
                setReviews(reviews.filter(r => r._id !== userReview._id))
            } catch (error) {
                console.log(error);
                toast.success(error.response.data, { autoClose: 500, theme: 'colored' })
            }
        } else {
            toast.success("Access denied", { autoClose: 500, theme: 'colored' })
        }
    }
    const sendEditResponse = async () => {
        if (!editComment && !value) {
            toast.error("Please Fill the all Fields", { autoClose: 500, })
        }
        else if (editComment.length <= 4) {
            toast.error("Please add more than 4 characters", { autoClose: 500, })
        }
        else if (value <= 0) {
            toast.error("Please add rating", { autoClose: 500, })
        }
        else if (editComment.length >= 4 && value > 0) {
            try {
                if (authToken) {
                    const response = await axios.put(`${process.env.REACT_APP_EDIT_REVIEW}`,
                        { id: userReview._id, comment: editComment, rating: value },
                        {
                            headers: {
                                'Authorization': `Bearer ${authToken}`
                            }
                        })
                    toast.success(response.data.msg, { autoClose: 500, })
                    fetchReviews()
                    setEdit(false)
                }
            }
            catch (error) {
                toast.error("Something went wrong", { autoClose: 600, })
            }
        }
    }
    return (
        <Grid container wrap="nowrap" spacing={2} sx={{ 
            backgroundColor: '#fff', 
            boxShadow: '0px 8px 13px rgba(0, 0, 0, 0.2)', 
            borderRadius: 5, 
            margin: '20px auto', 
            width: '100%', 
            height: 'auto',
            padding: 2
        }}>
            <Grid item>
                <Avatar alt="Customer Avatar" />
            </Grid>
            <Grid justifyContent="left" item xs zeroMinWidth >
                <h4 style={{ margin: 0, textAlign: "left" }}>
                    {userReview?.user?.firstName || 'Anonymous'} {userReview?.user?.lastName || ''}
                </h4>
                <p style={{ textAlign: "left", marginTop: 10 }}>
                    {!edit && <Rating name="read-only" value={Number(userReview.rating)} readOnly precision={0.5} />}
                    {edit &&
                        <Rating
                            name="simple-controlled"
                            value={value}
                            precision={0.5}
                            onChange={(event, newValue) => {
                                setValue(newValue);
                            }}
                        />}
                </p>
                <p style={{ textAlign: "left", wordBreak: 'break-word', paddingRight: 10, margin: '10px 0' }}>
                    {!edit && userReview.content}
                </p>
                {edit && (
                    <TextField
                        id="standard-basic"
                        value={editComment}
                        onChange={(e) => setEditComment(e.target.value)}
                        label="Edit Review"
                        multiline
                        className='comment'
                        variant="standard"
                        sx={{ width: '90%' }}
                    />
                )}

                {edit && (
                    <div style={{ display: 'flex', gap: 5, margin: 10 }}>
                        <Button 
                            sx={{ width: 10, borderRadius: '30px' }} 
                            variant='contained' 
                            onClick={sendEditResponse}
                        >
                            <AiOutlineSend />
                        </Button>
                        <Button 
                            sx={{ width: 10, borderRadius: '30px' }} 
                            variant='contained' 
                            color='error' 
                            onClick={() => setEdit(false)}
                        >
                            <GiCancel style={{ fontSize: 15, color: 'white' }} />
                        </Button>
                    </div>
                )}

                <p style={{ textAlign: "left", color: "gray", margin: "20px 0" }}>
                    {formatDate(userReview.createdAt)} {formatTime(userReview.createdAt)}
                </p>

                {(authUser === userReview?.user?.email || isAdmin) && (
                    <Box sx={{ height: 20, transform: 'translateZ(0px)', flexGrow: 1 }}>
                        <SpeedDial
                            ariaLabel="SpeedDial basic example"
                            sx={{ position: 'absolute', bottom: 16, right: 16 }}
                            icon={<SpeedDialIcon />}
                        >
                            <SpeedDialAction
                                icon={<AiFillEdit />}
                                tooltipTitle={"Edit"}
                                onClick={() => setEdit(true)}
                            />
                            <SpeedDialAction
                                icon={<AiFillDelete />}
                                tooltipTitle={"Delete"}
                                onClick={isAdmin ? deleteCommentByAdmin : handleDeleteComment}
                            />
                        </SpeedDial>
                    </Box>
                )}
            </Grid>
        </Grid>
    )
}

export default CommentCard