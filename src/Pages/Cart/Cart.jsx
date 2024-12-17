import React, { useContext, useEffect, useState } from 'react'
import { ContextFunction } from '../../Context/Context';
import {
    Button, Typography, Dialog, DialogActions, DialogContent, 
    Container, CssBaseline, Box
} from '@mui/material'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { toast } from 'react-toastify'
import { AiFillCloseCircle, AiOutlineLogin } from 'react-icons/ai'
import ProductCard from '../../Components/Card/Product Card/ProductCard'
import OrderSummary from './OrderSummary';
import { EmptyCart } from '../../Assets/Images/Image';
import { Transition, clearAuth } from '../../Constants/Constant';
import CopyRight from '../../Components/CopyRight/CopyRight';

const Cart = () => {
    const { cart, setCart } = useContext(ContextFunction)
    const [loading, setLoading] = useState(true)
    const [openAlert, setOpenAlert] = useState(false)
    const shippingCost = 100

    const navigate = useNavigate()
    const authToken = localStorage.getItem('Authorization')
    const setProceed = authToken ? true : false

    useEffect(() => {
        if (setProceed) {
            getCart()
        } else {
            setLoading(false)
            setOpenAlert(true)
        }
        window.scroll(0, 0)
    }, [setProceed])

    const getCart = async () => {
        try {
            const response = await axios.get(
                `${process.env.REACT_APP_GET_CART}`,
                {
                    headers: {
                        'Authorization': `Bearer ${authToken}`,
                        'Content-Type': 'application/json'
                    }
                }
            );
            
            if (response.data) {
                setCart(response.data);
            }
            setLoading(false);
        } catch (error) {
            if (error.response?.status === 401) {
                toast.error("Session expired. Please login again", {
                    autoClose: 500,
                    theme: 'colored'
                });
                clearAuth();
                navigate('/login');
            } else {
                console.error('Cart Error:', error);
                toast.error("Error loading cart", {
                    autoClose: 500,
                    theme: 'colored'
                });
            }
            setLoading(false);
        }
    }

    const removeFromCart = async (productId) => {
        try {
            await axios.delete(
                `${process.env.REACT_APP_DELETE_CART}`,
                {
                    headers: {
                        'Authorization': `Bearer ${authToken}`
                    },
                    data: { productid: productId }
                }
            )
            
            const updatedItems = cart.Items.filter(item => item.ProductID !== productId);
            
            const newTotal = updatedItems.reduce((total, item) => {
                return total + (item.Price * item.Quantity)
            }, 0);

            const updatedCart = {
                ...cart,
                Items: updatedItems,
                Totals: newTotal
            }
            
            setCart(updatedCart)
            toast.success('Removed from cart')
        } catch (error) {
            toast.error(error.response?.data?.error || 'Failed to remove item')
        }
    }

    const proceedToCheckout = () => {
        if (!cart.Items || cart.Items.length === 0) {
            toast.error('Please add items to cart first')
            return
        }
        sessionStorage.setItem('totalAmount', cart.Totals)
        navigate('/checkout')
    }

    const updateQuantity = async (productId, newQuantity) => {
        try {
            console.log('Updating quantity:', {
                productId,
                newQuantity,
                requestBody: { 
                    productId: productId,
                    newQuantity: newQuantity 
                }
            });

            const response = await axios.put(
                `${process.env.REACT_APP_GET_CART}/quantity`,
                { 
                    productId: productId,
                    newQuantity: newQuantity 
                },
                {
                    headers: {
                        'Authorization': `Bearer ${authToken}`,
                        'Content-Type': 'application/json'
                    }
                }
            );

            console.log('Update response:', response.data);

            if (response.data) {
                setCart(response.data);
            } else {
                const updatedCart = {
                    ...cart,
                    Items: cart.Items.map(item => 
                        item.ProductID === productId 
                            ? { ...item, Quantity: newQuantity }
                            : item
                    )
                };
                
                updatedCart.Totals = updatedCart.Items.reduce((total, item) => {
                    return total + (item.Price * item.Quantity);
                }, 0);

                setCart(updatedCart);
            }
        } catch (error) {
            console.error('Update error:', {
                status: error.response?.status,
                data: error.response?.data,
                error: error.message
            });
            toast.error(error.response?.data?.error || 'Failed to update quantity');
        }
    };

    if (loading) return <div>Loading...</div>

    return (
        <>
            <CssBaseline />
            <Container fixed maxWidth>
                <Typography variant='h3' sx={{ textAlign: 'center', mt: 10, color: '#1976d2', fontWeight: 'bold' }}>
                    Cart
                </Typography>

                <Container sx={{ display: 'flex', flexDirection: "column", mb: 10 }}>
                    {cart?.Items && cart.Items.length > 0 ? (
                        <>
                            <Box sx={{ display: 'flex', justifyContent: 'center', flexWrap: 'wrap', gap: 2 }}>
                                {cart.Items.map(item => (
                                    <ProductCard 
                                        key={item.ProductID}
                                        prod={{
                                            ID: item.ProductID,
                                            Name: item.Name,
                                            Price: item.Price,
                                            Image: item.Image,
                                            Quantity: item.Quantity,
                                            Rating: item.Rating
                                        }}
                                        onRemove={() => removeFromCart(item.ProductID)}
                                        onUpdateQuantity={updateQuantity}
                                        showControls={true}
                                    />
                                ))}
                            </Box>

                            <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
                                <OrderSummary 
                                    proceedToCheckout={proceedToCheckout}
                                    total={cart.Totals}
                                    shippingCost={shippingCost}
                                />
                            </Box>
                        </>
                    ) : (
                        <Box sx={{ width: '100%', display: 'flex', justifyContent: 'center' }}>
                            <div className="main-card">
                                <img src={EmptyCart} alt="Empty_cart" className="empty-cart-img" />
                                <Typography variant='h6' sx={{ textAlign: 'center', color: '#1976d2' }}>
                                    Your Cart is Empty
                                </Typography>
                            </div>
                        </Box>
                    )}
                </Container>
            </Container>

            <Dialog
                open={openAlert}
                TransitionComponent={Transition}
                keepMounted
                onClose={() => navigate('/')}
            >
                <DialogContent sx={{ width: { xs: 280, md: 350, xl: 400 }, display: 'flex', justifyContent: 'center' }}>
                    <Typography variant='h5'>Please Login To Proceed</Typography>
                </DialogContent>
                <DialogActions sx={{ display: 'flex', justifyContent: 'space-evenly' }}>
                    <Button 
                        variant='contained' 
                        onClick={() => navigate('/login')} 
                        endIcon={<AiOutlineLogin />} 
                        color='primary'
                    >
                        Login
                    </Button>
                    <Button 
                        variant='contained' 
                        color='error' 
                        endIcon={<AiFillCloseCircle />} 
                        onClick={() => navigate('/')}
                    >
                        Close
                    </Button>
                </DialogActions>
            </Dialog>
            
            <CopyRight sx={{ mt: 8, mb: 10 }} />
        </>
    )
}

export default Cart