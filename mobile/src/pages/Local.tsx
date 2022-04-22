import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Local = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Local Product"/>
        <Container>
            <View>
                <Text>Local Product</Text>
            </View>
        </Container>
        </>
    );
};

export default Local;