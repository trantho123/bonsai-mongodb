import { Slide } from "@mui/material";
import axios from "axios";
import { forwardRef } from "react";
const getCart = async (setProceed, setCart, authToken) => {
    if (setProceed) {
        const { data } = await axios.get(`${process.env.REACT_APP_GET_CART}`,
            {
                headers: {
                    'Authorization': authToken
                }
            })
        setCart(data);
    }
}
const getWishList = async (setProceed, setWishlistData, authToken) => {
    if (setProceed) {
        const { data } = await axios.get(`${process.env.REACT_APP_GET_WISHLIST}`,
            {
                headers: {
                    'Authorization': authToken
                }
            })
        setWishlistData(data)
    }
}
const handleLogOut = (setProceed, toast, navigate, setOpenAlert) => {
    if (setProceed) {
        try {
            // Xóa token và các dữ liệu auth
            clearAuth();
            
            // Hiển thị thông báo thành công
            toast.success("Logout Successfully", { 
                autoClose: 500, 
                theme: 'colored' 
            });
            
            // Đóng dialog nếu đang mở
            if (setOpenAlert) {
                setOpenAlert(false);
            }
            
            // Chuyển về trang chủ
            navigate('/');
            
        } catch (error) {
            console.error("Logout error:", error);
            toast.error("Error during logout", { 
                autoClose: 500, 
                theme: 'colored' 
            });
        }
    } else {
        toast.error("User is already logged out", { 
            autoClose: 500, 
            theme: 'colored' 
        });
    }
}

const handleClickOpen = (setOpenAlert) => {
    setOpenAlert(true);
};

const handleClose = (setOpenAlert) => {
    setOpenAlert(false);
};
const getAllProducts = async (setData) => {
    try {
        const { data } = await axios.get(process.env.REACT_APP_FETCH_PRODUCT);
        setData(data)


    } catch (error) {
        console.log(error);
    }
}

const getSingleProduct = async (setProduct, id, setLoading) => {
    try {
        console.log('Calling API:', `${process.env.REACT_APP_FETCH_PRODUCT_DETAIL}/${id}`);
        
        const { data } = await axios.get(`${process.env.REACT_APP_FETCH_PRODUCT_DETAIL}/${id}`)
        
        console.log('API Response:', {
            url: `${process.env.REACT_APP_FETCH_PRODUCT_DETAIL}/${id}`,
            status: 'success',
            data: data
        });
        
        setProduct(data)
        setLoading(false);
    } catch (error) {
        console.error('API Error:', {
            url: `${process.env.REACT_APP_FETCH_PRODUCT_DETAIL}/${id}`,
            status: error.response?.status,
            error: error.response?.data || error.message
        });
        setLoading(false);
    }
}

const Transition = forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export const clearAuth = () => {
    // Xóa token và các dữ liệu auth khác
    localStorage.removeItem('Authorization');
    localStorage.removeItem('isAdmin');
    
    // Xóa các dữ liệu trong sessionStorage nếu có
    sessionStorage.clear();
}

export const CURRENCY_SYMBOL = 'VND';

export { getCart, getWishList, handleClickOpen, handleClose, handleLogOut, getAllProducts, getSingleProduct, Transition }