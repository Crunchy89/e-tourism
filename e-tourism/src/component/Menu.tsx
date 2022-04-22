import React from 'react';
import {StyleSheet,FlatList} from 'react-native';
import LinearGradient from 'react-native-linear-gradient';
import ButtonIcon from './ButtonIcon';
import { menu } from './data';


const Menu  = (props) => {
    return (
        <LinearGradient start={{x: 0, y: 0}} end={{x: 1, y: 0}} colors={['rgba(255,226,89,.8)', 'rgba(255,204,86,.8)', 'rgba(255,167,81,.8)']} style={styles.menuContainer}>
            <FlatList
                numColumns={4}
                keyExtractor={(_, index) => index.toString()}
                data={menu} 
                renderItem={({item}) =>  (
                        <ButtonIcon
                        title={item.title}
                        image={item.image}
                        link={item.link}
                        navigation={props.navigation}/>
                    )
                }
                columnWrapperStyle={styles.flat}
                showsVerticalScrollIndicator={false}
                showsHorizontalScrollIndicator={false}
            />
        </LinearGradient>
    );
};

const styles = StyleSheet.create({
    menuContainer:{
        flex:3,
        borderRadius:5,
        marginTop:10,
        marginBottom:10,
        padding: 10,
    },
    flat:{
        flexWrap:"wrap",
        justifyContent:"space-around",
    }
  });

export default Menu;