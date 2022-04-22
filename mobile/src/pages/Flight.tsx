import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Flight = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Flight Plan"/>
        <Container>
            <View>
                <Text>Flight Plan</Text>
            </View>
        </Container>
        </>
    );
};

export default Flight;