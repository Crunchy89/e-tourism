import React from 'react';
import {View,Text,StyleSheet,Button} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';


const Market = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Market"/>
        <Container>
            <View>
                <Text>Market</Text>
            </View>
        </Container>
        </>
    );
};

export default Market;