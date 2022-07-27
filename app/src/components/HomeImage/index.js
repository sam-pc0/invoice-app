import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';


export default function HomaImage() {
    return (
        <Paper
            style={{ margin: '20px' }}
            sx={{
                position: 'relative',
                backgroundColor: 'grey.900',
                color: '#fff',
                textAlign: 'center',
                height: "400px",
                backgroundSize: 'cover',
                backgroundRepeat: 'no-repeat',
                backgroundPosition: 'center',
                backgroundImage: `url('https://www.eluniverso.com/resizer/vBJTf88IaqzzC1UlJ86dsHZH3Bo=/800x533/smart/filters:quality(70)/cloudfront-us-east-1.images.arcpublishing.com/eluniverso/525O7HO53NCTXEGD7KBFTZLTCQ.jpg')`,
            }}
        >
            {/* Increase the priority of the hero background image */}
            {<img style={{ display: 'none' }} src={'https://www.eluniverso.com/resizer/vBJTf88IaqzzC1UlJ86dsHZH3Bo=/800x533/smart/filters:quality(70)/cloudfront-us-east-1.images.arcpublishing.com/eluniverso/525O7HO53NCTXEGD7KBFTZLTCQ.jpg'} alt={'main image description'} />}
            <Box
                sx={{
                    position: 'absolute',
                    top: 0,
                    bottom: 0,
                    right: 0,
                    left: 0,
                    backgroundColor: 'rgba(0,0,0,.6)',
                }}
            />
            <Grid container style={{ paddingLeft: '450px' }}>
                <Grid item md={6}>
                    <Box
                        sx={{
                            position: 'relative',
                            p: { xs: 19, md: 15 },
                            pr: { md: 0 },
                        }}
                    >
                        <Typography component="h1" variant="h3" color="inherit" gutterBottom>
                            Busca tus productos favoritos a los mejores precios
                        </Typography>

                    </Box>
                </Grid>
            </Grid>
        </Paper>
    )
}