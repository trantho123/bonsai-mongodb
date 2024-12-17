import { 
    Button, 
    Card, 
    CardContent, 
    CardMedia, 
    Rating, 
    Tooltip, 
    Typography,
    Box,
    CardActions
} from '@mui/material'
import { AiFillDelete } from 'react-icons/ai'
import styles from './CartCard.module.css'

const CartCard = ({ product, removeFromCart }) => {
    console.log('CartCard received product:', product)

    return (
        <Card className={styles.main_cart}>
            <Box className={styles.img_box}>
                <CardMedia
                    component="img"
                    image={product.Image}
                    alt={product.Name}
                    className={styles.img}
                />
            </Box>

            <CardContent>
                <Typography variant="h6" sx={{ textAlign: "center" }}>
                    {product.Name?.length > 20 ? product.Name.slice(0, 20) + '...' : product.Name}
                </Typography>

                <Box sx={{ 
                    display: 'flex', 
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    mt: 2
                }}>
                    <Typography variant="body1">
                        Quantity: {product.Quantity}
                    </Typography>
                    <Typography variant="h6" color="primary">
                        VND {product.Price}
                    </Typography>
                </Box>

                <Box sx={{ mt: 1 }}>
                    <Typography variant="body2" color="text.secondary">
                        Total: VND {product.Price * product.Quantity}
                    </Typography>
                </Box>
            </CardContent>

            <CardActions sx={{ justifyContent: 'space-between', px: 2 }}>
                <Tooltip title="Remove From Cart">
                    <Button
                        variant="contained"
                        color="error"
                        onClick={() => removeFromCart(product.productId)}
                        startIcon={<AiFillDelete />}
                        size="small"
                    >
                        Remove
                    </Button>
                </Tooltip>
            </CardActions>
        </Card>
    )
}

export default CartCard