import { 
    Card, 
    CardActionArea, 
    CardActions, 
    Rating, 
    CardContent, 
    Typography,
    Button,
    ButtonGroup,
    Box,
    IconButton,
    Tooltip
} from '@mui/material';
import { AiFillDelete } from 'react-icons/ai';
import { BiMinus, BiPlus } from 'react-icons/bi';
import styles from './ProductCard.module.css'

export default function ProductCard({ prod, onRemove, onUpdateQuantity, showControls = false }) {
    if (!prod) {
        return null;
    }

    const handleIncreaseQuantity = () => {
        console.log('Increase clicked:', {
            productId: prod.ID,
            currentQuantity: prod.Quantity,
            newQuantity: prod.Quantity + 1
        });
        onUpdateQuantity && onUpdateQuantity(prod.ID, prod.Quantity + 1);
    };

    const handleDecreaseQuantity = () => {
        if (prod.Quantity > 1) {
            console.log('Decrease clicked:', {
                productId: prod.ID,
                currentQuantity: prod.Quantity,
                newQuantity: prod.Quantity - 1
            });
            onUpdateQuantity && onUpdateQuantity(prod.ID, prod.Quantity - 1);
        }
    };

    return (
        <Card className={styles.main_card}>
            {showControls && onRemove && (
                <Box sx={{ 
                    position: 'absolute', 
                    top: 8, 
                    right: 8, 
                    zIndex: 1 
                }}>
                    <Tooltip title="Remove from cart">
                        <IconButton
                            color="error"
                            onClick={() => onRemove(prod.ID)}
                            size="small"
                            sx={{ 
                                bgcolor: 'rgba(255,255,255,0.9)',
                                '&:hover': {
                                    bgcolor: 'rgba(255,255,255,1)'
                                }
                            }}
                        >
                            <AiFillDelete />
                        </IconButton>
                    </Tooltip>
                </Box>
            )}

            <CardActionArea className={styles.card_action}>
                <Box className={styles.cart_box}>
                    <img 
                        alt={prod.Name || 'Product'} 
                        src={prod.Image} 
                        loading='lazy' 
                        className={styles.cart_img} 
                    />
                </Box>
                <CardContent>
                    <Typography 
                        gutterBottom 
                        variant="h6" 
                        sx={{ textAlign: "center" }}
                    >
                        {prod.Name ? 
                            (prod.Name.length > 20 ? prod.Name.slice(0, 20) + '...' : prod.Name)
                            : 'Product Name'
                        }
                    </Typography>
                    
                    <Rating 
                        name="read-only" 
                        value={Number(prod.Rating) || 0}
                        precision={0.5}
                        readOnly 
                    />
                    
                    <Typography variant="h6" color="primary" sx={{ mt: 1 }}>
                        VND {prod.Price || 0}
                    </Typography>

                    {showControls && (
                        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
                            <ButtonGroup size="small" variant="outlined">
                                <IconButton 
                                    onClick={handleDecreaseQuantity}
                                    disabled={!prod.Quantity || prod.Quantity <= 1}
                                >
                                    <BiMinus />
                                </IconButton>
                                <Button disabled>{prod.Quantity || 0}</Button>
                                <IconButton 
                                    onClick={handleIncreaseQuantity}
                                    disabled={!prod.Quantity}
                                >
                                    <BiPlus />
                                </IconButton>
                            </ButtonGroup>
                        </Box>
                    )}
                </CardContent>
            </CardActionArea>
        </Card>
    );
}