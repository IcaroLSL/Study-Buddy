import Ionicons from '@expo/vector-icons/Ionicons';
import { View, Pressable, Text } from 'react-native';

interface NotificationCardProps {
  id: string | number;
  title: string;
  time: string;
  description?: string;
  onDelete?: (id: string | number) => void;
}

export default function NotificationCard({
  id,
  title,
  time,
  description,
  onDelete,
}: NotificationCardProps) {
  return (
    <View className="flex-row items-center rounded-lg border border-gray-200 bg-white p-3 dark:border-gray-700 dark:bg-gray-800">
      <View className="mr-3 rounded-full bg-primary-500 px-2 py-4">
        <Ionicons name="calendar-clear" size={20} color="white" />
      </View>

      <View className="flex-1">
        <Text className="font-semibold text-lg text-primary-600 dark:text-primary-400">
          {title}
        </Text>

        <Text className="text-md mb-1 text-gray-600 dark:text-gray-300">
          {time} - {description || 'Sem descrição'}
        </Text>
      </View>

      <Pressable
        className="rounded p-1 active:bg-gray-200 dark:active:bg-gray-600"
        onPress={() => onDelete?.(id)}
      >
        {({ pressed }) => (
            <Ionicons
            name="close"
            size={16}
            color={pressed ? '#EF4444' : '#9CA3AF'}
            />
        )}
      </Pressable>
    </View>
  );
}