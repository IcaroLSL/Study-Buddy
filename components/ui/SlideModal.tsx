import React, { useEffect, useRef, useCallback, useState } from 'react';
import { Animated, Modal, Pressable, Text, View, useWindowDimensions } from 'react-native';
import Ionicons from '@expo/vector-icons/Ionicons';

interface ModalProps {
  visible: boolean;
  title?: string;
  onClose: () => void;
  children: React.ReactNode;
}

export default function SlideModalComponent({
  visible,
  title,
  onClose,
  children,
}: ModalProps) {
  const { width } = useWindowDimensions();
  const translateX = useRef(new Animated.Value(width)).current;
  const backdropOpacity = useRef(new Animated.Value(0)).current;
  const [internalVisible, setInternalVisible] = useState(false);

  const runCloseAnimation = useCallback(() => {
    Animated.parallel([
      Animated.timing(translateX, {
        toValue: width,
        duration: 250,
        useNativeDriver: true,
      }),
      Animated.timing(backdropOpacity, {
        toValue: 0,
        duration: 200,
        useNativeDriver: true,
      }),
    ]).start(() => {
      setInternalVisible(false);
    });
  }, [width]);

  // Sync with external visible prop
  useEffect(() => {
    if (visible && !internalVisible) {
      // Opening: reset values and show modal
      translateX.setValue(width);
      backdropOpacity.setValue(0);
      setInternalVisible(true);
    } else if (!visible && internalVisible) {
      // External close requested: animate out
      runCloseAnimation();
    }
  }, [visible, internalVisible, width, runCloseAnimation]);

  // Watches internalVisible to trigger open animation after mount
  useEffect(() => {
    if (internalVisible) {
      Animated.parallel([
        Animated.timing(translateX, {
          toValue: 0,
          duration: 280,
          useNativeDriver: true,
        }),
        Animated.timing(backdropOpacity, {
          toValue: 1,
          duration: 200,
          useNativeDriver: true,
        }),
      ]).start();
    }
  }, [internalVisible]);

  const handleClose = useCallback(() => {
    runCloseAnimation();
    // Notify parent immediately
    onClose();
  }, [runCloseAnimation, onClose]);

  if (!internalVisible) {
    return null;
  }

  return (
    <Modal visible={true} transparent animationType="none" onRequestClose={handleClose} >
        <View className="flex-1">
            <Pressable className="absolute inset-0" onPress={handleClose} >
                <Animated.View pointerEvents="none" className="absolute inset-0 bg-black/50" style={{ opacity: backdropOpacity }} />
            </Pressable>
            <Animated.View className="ml-auto h-full w-4/5 max-w-sm bg-white p-6 dark:bg-gray-800" style={{ transform: [{ translateX }] }} >
                {/* Cabeçalho */}
                <View className="w-full flex-row items-center justify-between">
                    {title && (
                    <Text className="text-2xl font-bold text-gray-800 dark:text-white">
                        {title}
                    </Text>
                    )}

                    <Pressable onPress={handleClose}>
                      {({ pressed }) => (
                        <Ionicons
                          name="close"
                          size={28}
                          color={pressed ? '#EF4444' : '#9CA3AF'}
                        />
                      )}
                    </Pressable>
                </View>

                {/* Separador */}
                <View className="-mx-6 my-4 h-px bg-gray-300 dark:bg-gray-600" />

                {/* Children */}
                <View className="w-full">
                    {children}
                </View>
            </Animated.View>
        </View>
    </Modal>
  );
}
