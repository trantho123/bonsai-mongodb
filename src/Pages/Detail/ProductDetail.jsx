import './Productsimilar.css'
import React, { useEffect, useState, useContext } from 'react'
import { useParams, Link } from 'react-router-dom'
import {
    Box,
    Button,
    Container,
    Tooltip,
    Typography,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    Chip,
    Rating,
    ButtonGroup,
    Skeleton,
    IconButton,
} from '@mui/material';
import { MdAddShoppingCart } from 'react-icons/md'
import { AiFillHeart, AiFillCloseCircle, AiOutlineLogin, AiOutlineShareAlt } from 'react-icons/ai'
import { TbDiscount2 } from 'react-icons/tb'
import axios from 'axios';
import { toast } from 'react-toastify';
import { ContextFunction } from '../../Context/Context';
import ProductReview from '../../Components/Review/ProductReview';
import ProductCard from '../../Components/Card/Product Card/ProductCard';
import { Transition, getSingleProduct } from '../../Constants/Constant';
import CopyRight from '../../Components/CopyRight/CopyRight';



const ProductDetail = () => {
    const { cart, setCart, wishlistData, setWishlistData } = useContext(ContextFunction)
    const [openAlert, setOpenAlert] = useState(false);
    const { id, cat } = useParams()
    const [product, setProduct] = useState([])
    const [similarProduct, setSimilarProduct] = useState([])
    const [productQuantity, setProductQuantity] = useState(1)
    const [loading, setLoading] = useState(true);


    let authToken = localStorage.getItem('Authorization')
    let setProceed = authToken ? true : false


    useEffect(() => {
        const fetchData = async () => {
            try {
                setLoading(true);
                console.log('Fetching product:', `${process.env.REACT_APP_FETCH_PRODUCT_DETAIL}/${id}`);
                
                const response = await axios.get(`${process.env.REACT_APP_FETCH_PRODUCT_DETAIL}/${id}`);
                console.log('Product data:', response.data);
                
                const productData = response.data;
                setProduct(productData);
                setLoading(false);

                if (productData.tags && productData.tags.length > 0) {
                    console.log('Fetching similar products for tags:', productData.tags);
                    const tagNames = productData.tags.map(tag => tag.Name);
                    const { data } = await axios.post(
                        `${process.env.REACT_APP_PRODUCT_TYPE}`, 
                        { tags: tagNames }
                    );
                    const filtered = data.filter(p => p.id !== id);
                    setSimilarProduct(filtered);
                }
            } catch (error) {
                console.error("Error fetching data:", error.response || error);
                setLoading(false);
                toast.error(error.response?.data?.message || "Failed to load product details");
            }
        };

        if (id) {
            fetchData();
            window.scroll(0, 0);
        }
    }, [id]);

    if (loading) {
        return (
            <Container>
                <Box sx={{ mt: 3 }}>
                    <Skeleton variant="rectangular" height={400} />
                    <Skeleton variant="text" sx={{ mt: 2 }} />
                    <Skeleton variant="text" />
                    <Skeleton variant="text" width="60%" />
                </Box>
            </Container>
        );
    }

    if (!product || Object.keys(product).length === 0) {
        return (
            <Container>
                <Box sx={{ mt: 3, textAlign: 'center' }}>
                    <Typography variant="h5" color="error">
                        Product not found or failed to load
                    </Typography>
                </Box>
            </Container>
        );
    }

    const addToCart = async (product) => {
        if (setProceed) {
            try {
                const { data } = await axios.post(
                    `${process.env.REACT_APP_ADD_CART}`, 
                    { 
                        productId: product.id, 
                        quantity: productQuantity 
                    }, 
                    {
                        headers: {
                            'Authorization': `Bearer ${authToken}`,
                            'Content-Type': 'application/json'
                        }
                    }
                );
                setCart(data);
                toast.success("Added To Cart", { autoClose: 500, theme: 'colored' });
            } catch (error) {
                toast.error(error.response?.data?.error || "Failed to add to cart", { 
                    autoClose: 500, 
                    theme: 'colored' 
                });
            }
        } else {
            setOpenAlert(true);
        }
    }
    const addToWhishList = async (product) => {
        if (setProceed) {
            try {
                const { data } = await axios.post(`${process.env.REACT_APP_ADD_WISHLIST}`, { _id: product._id }, {
                    headers: {
                        'Authorization': authToken
                    }
                })
                setWishlistData(data)
                setWishlistData([...wishlistData, product])
                toast.success("Added To Wishlist", { autoClose: 500, theme: 'colored' })
            }
            catch (error) {
                toast.error(error.response.data.msg, { autoClose: 500, theme: 'colored' })
            }
        }
        else {
            setOpenAlert(true);
        }

    };
    const shareProduct = (product) => {

        const data = {
            text: product.name,
            title: "e-shopit",
            url: `https://e-shopit.vercel.app/Detail/type/${cat}/${id}`
        }
        if (navigator.canShare && navigator.canShare(data)) {
            navigator.share(data);
        }
        else {
            toast.error("browser not support", { autoClose: 500, theme: 'colored' })
        }

    }
    const getSimilarProducts = async () => {
        try {
            // Lấy tags từ sản phẩm hiện tại
            const tagNames = product.tags.map(tag => tag.Name);
            
            // Gọi API với mảng tags
            const { data } = await axios.post(
                `${process.env.REACT_APP_PRODUCT_TYPE}`, 
                { tags: tagNames }
            );

            // Lọc bỏ sản phẩm hiện tại khỏi kết quả
            const filtered = data.filter(p => p.id !== id);
            
            setSimilarProduct(filtered);
        } catch (error) {
            console.error("Error fetching similar products:", error);
        }
    }
    const increaseQuantity = () => {
        setProductQuantity((prev) => prev + 1)
        if (productQuantity >= 5) {
            setProductQuantity(5)
        }
    }
    const decreaseQuantity = () => {
        setProductQuantity((prev) => prev - 1)
        if (productQuantity <= 1) {
            setProductQuantity(1)
        }
    }
    return (
        <>
            <Container maxWidth='xl' >
                <Dialog
                    open={openAlert}
                    TransitionComponent={Transition}
                    keepMounted
                    onClose={() => setOpenAlert(false)}
                    aria-describedby="alert-dialog-slide-description"
                >
                    <DialogContent sx={{ width: { xs: 280, md: 350, xl: 400 } }}>
                        <DialogContentText style={{ textAlign: 'center' }} id="alert-dialog-slide-description">
                            Please Login To Proceed
                        </DialogContentText>
                    </DialogContent>
                    <DialogActions sx={{ display: 'flex', justifyContent: 'space-evenly' }}>
                        <Link to="/login"> <Button variant='contained' endIcon={<AiOutlineLogin />} color='primary'>Login</Button></Link>
                        <Button variant='contained' color='error'
                            onClick={() => setOpenAlert(false)} endIcon={<AiFillCloseCircle />}>Close</Button>
                    </DialogActions>
                </Dialog>

                <main className='main-content'>
                    <div className="product-image">
                        <div className='detail-img-box'  >
                            <img alt={product.name} src={product.image} className='detail-img' />
                            <br />
                        </div>
                    </div>
                    <section className='product-details'>
                        <Typography variant='h4'>{product.name}</Typography>

                        <Typography >
                            {product.description}
                        </Typography>
                        <Typography>
                            <div className="chip">
                                {product.tags && product.tags.map((tag, index) => (
                                    <Chip 
                                        key={tag.ID} 
                                        label={tag.Name}
                                        variant="outlined"
                                    />
                                ))}
                            </div>
                        </Typography>
                        <Chip
                            label={product.price > 1000 ? "Upto 9% off" : "Upto 38% off"}
                            variant="outlined"
                            sx={{ background: '#1976d2', color: 'white', width: '150px', fontWeight: 'bold' }}
                            avatar={<TbDiscount2 color='white' />}


                        />
                        <div style={{ display: 'flex', gap: 20 }}>
                            <Typography variant="h6" color="red"> VND {product.price > 1000 ? product.price + 1000 : product.price + 300} </Typography>
                            <Typography variant="h6" color="primary">
                                VND {product.price}
                            </Typography>
                        </div>
                        <Box
                            sx={{
                                display: 'flex',
                                flexDirection: 'column',
                                // background: 'red',
                                '& > *': {
                                    m: 1,
                                },
                            }}
                        >
                            <ButtonGroup variant="outlined" aria-label="outlined button group">
                                <Button onClick={increaseQuantity}>+</Button>
                                <Button>{productQuantity}</Button>
                                <Button onClick={decreaseQuantity}>-</Button>
                            </ButtonGroup>
                        </Box>
                        <Rating name="read-only" value={Math.round(product.rating)} readOnly precision={0.5} />
                        <div style={{ display: 'flex' }} >
                            <Tooltip title='Add To Cart'>
                                <Button variant='contained' className='all-btn' startIcon={<MdAddShoppingCart />} onClick={(() => addToCart(product))}>Buy</Button>
                            </Tooltip>
                            <Tooltip title='Add To Wishlist'>
                                <Button style={{ marginLeft: 10, }} size='small' variant='contained' className='all-btn' onClick={(() => addToWhishList(product))}>
                                    {<AiFillHeart fontSize={21}/>}
                                </Button>

                            </Tooltip>
                            <Tooltip title='Share'>
                                <Button style={{ marginLeft: 10 }} variant='contained' className='all-btn' startIcon={<AiOutlineShareAlt />} onClick={() => shareProduct(product)}>Share</Button>
                            </Tooltip>

                        </div>
                    </section>
                </main>
                <ProductReview setProceed={setProceed} authToken={authToken} id={id} setOpenAlert={setOpenAlert} />


                <Typography sx={{ marginTop: 10, marginBottom: 5, fontWeight: 'bold', textAlign: 'center' }}>Similar Products</Typography>
                <Box>
                    <Box className='similarProduct' sx={{ display: 'flex', overflowX: 'auto', marginBottom: 10 }}>
                        {similarProduct && similarProduct.length > 0 && similarProduct
                            .filter(prod => prod.id !== id)
                            .map(prod => (
                                <Link to={`/product/${prod.id}`} key={prod.id}>
                                    <ProductCard 
                                        prod={{
                                            _id: prod.id,
                                            Name: prod.name,
                                            Price: prod.price,
                                            Image: prod.image,
                                            Description: prod.description,
                                            Rating: prod.rating,
                                            Tags: prod.tags
                                        }} 
                                    />
                                </Link>
                            ))
                        }
                    </Box>
                </Box>

            </Container >
            <CopyRight   sx={{ mt: 8, mb: 10 }} />

        </>
    )
}

export default ProductDetail