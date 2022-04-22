import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Rental = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Rental"/>
        <Container>
            <View>
                <Text>Rental</Text>
            </View>
        </Container>
        </>
    );
};

export default Rental;