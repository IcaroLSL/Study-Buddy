import FontAwesome6 from '@expo/vector-icons/FontAwesome6';
import Ionicons from '@expo/vector-icons/Ionicons';
import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import React, { type ComponentProps, useMemo } from 'react';
import { Pressable, ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

type TabId = 'dashboard' | 'techniques' | 'community' | 'focus' | 'config';

type BottomTab = {
  id: TabId;
  label: string;
  icon: ComponentProps<typeof FontAwesome6>['name'];
};

interface ScreenContentProps {
  title: string;
  path: string;
  children?: React.ReactNode;
  activeTab?: TabId;
  notificationCount?: number;
  isDark?: boolean;
  onHomePress?: () => void;
  onNotificationsPress?: () => void;
  onTabChange?: (tabId: TabId) => void;
  onThemeToggle?: () => void;
}

const bottomTabs: BottomTab[] = [
  { id: 'dashboard', label: 'Início', icon: 'house' },
  { id: 'techniques', label: 'Técnicas', icon: 'lightbulb' },
  { id: 'community', label: 'Comunidade', icon: 'users' },
  { id: 'focus', label: 'Foco', icon: 'bolt' },
  { id: 'config', label: 'Configurações', icon: 'gear' },
];

export const ScreenContent: React.FC<ScreenContentProps> = ({
  title,
  path,
  children,
  activeTab,
  notificationCount = 0,
  isDark = false,
  onHomePress,
  onNotificationsPress,
  onTabChange,
  onThemeToggle,
}) => {
  const selectedTab = useMemo(
    () => activeTab ?? getTabFromRoute(path, title),
    [activeTab, path, title]
  );
  const backgroundClassName = isDark ? 'flex-1 bg-gray-900' : 'flex-1 bg-gray-50';
  const headerClassName = isDark
    ? 'border-b border-gray-700 bg-gray-800 shadow-sm'
    : 'border-b border-gray-200 bg-white shadow-sm';
  const textClassName = isDark ? 'text-gray-100' : 'text-gray-800';
  const mutedTextClassName = isDark ? 'text-gray-400' : 'text-gray-500';
  const iconColor = isDark ? '#f3f4f6' : '#1f2937';

  return (
    <SafeAreaView className={backgroundClassName}>
      <View className="mx-auto w-full max-w-md flex-1">
        <View className={headerClassName}>
          <View className="flex-row items-center justify-between px-4 py-3">
            <Pressable
              accessibilityLabel="Ir para o início"
              className="flex-row items-center gap-2"
              onPress={onHomePress}>
              <FontAwesome6 color="#0ea5e9" name="graduation-cap" size={24} />
              <Text className={`text-xl font-bold ${textClassName}`}>StudyBuddy</Text>
            </Pressable>

            <View className="flex-row items-center gap-3">
              <Pressable
                accessibilityLabel="Abrir notificações"
                className={styles.headerIconButton(isDark)}
                onPress={onNotificationsPress}>
                <Ionicons color={iconColor} name="notifications" size={20} />
                {notificationCount > 0 ? (
                  <View className="absolute right-0 top-0 h-4 min-w-4 items-center justify-center rounded-full bg-red-500 px-1">
                    <Text className="text-[10px] font-bold text-white">
                      {notificationCount > 9 ? '9+' : notificationCount}
                    </Text>
                  </View>
                ) : null}
              </Pressable>

              <Pressable
                accessibilityLabel="Alternar tema"
                className={styles.headerIconButton(isDark)}
                onPress={onThemeToggle}>
                <MaterialIcons
                  color={iconColor}
                  name={isDark ? 'light-mode' : 'dark-mode'}
                  size={20}
                />
              </Pressable>
            </View>
          </View>
        </View>

        <ScrollView
          className="flex-1"
          contentContainerClassName="px-4 py-4 pb-28"
          showsVerticalScrollIndicator={false}>
          <View className="mb-4 flex-row items-center justify-between">
            <Text className={`text-lg font-semibold ${textClassName}`}>{title}</Text>
            <Text className={`text-xs ${mutedTextClassName}`}>{path}</Text>
          </View>
          {children}
        </ScrollView>

        <View className={styles.bottomNav(isDark)}>
          <View className="mx-auto w-full max-w-md flex-row justify-around">
            {bottomTabs.map((tab) => {
              const isSelected = tab.id === selectedTab;
              const tabColor = isSelected ? '#0ea5e9' : isDark ? '#9ca3af' : '#6b7280';

              return (
                <Pressable
                  accessibilityLabel={`Abrir aba ${tab.label}`}
                  accessibilityRole="tab"
                  accessibilityState={{ selected: isSelected }}
                  className="min-w-[64px] items-center px-2 py-3"
                  key={tab.id}
                  onPress={() => onTabChange?.(tab.id)}>
                  <FontAwesome6 color={tabColor} name={tab.icon} size={20} />
                  <Text className="mt-1 text-xs font-medium" style={{ color: tabColor }}>
                    {tab.label}
                  </Text>
                  <View className={isSelected ? styles.activeTabIndicator : styles.tabIndicator} />
                </Pressable>
              );
            })}
          </View>
        </View>
      </View>
    </SafeAreaView>
  );
};

const getTabFromRoute = (path: string, title: string): TabId => {
  const route = `${path} ${title}`.toLowerCase();

  if (route.includes('technique') || route.includes('técnica')) return 'techniques';
  if (route.includes('community') || route.includes('comunidade')) return 'community';
  if (route.includes('focus') || route.includes('foco')) return 'focus';
  if (route.includes('config') || route.includes('settings')) return 'config';

  return 'dashboard';
};

const styles = {
  activeTabIndicator: 'mt-2 h-[3px] w-8 rounded-full bg-sky-500',
  bottomNav: (isDark: boolean) =>
    isDark
      ? 'absolute inset-x-0 bottom-0 border-t border-gray-700 bg-gray-800 shadow-lg'
      : 'absolute inset-x-0 bottom-0 border-t border-gray-200 bg-white shadow-lg',
  headerIconButton: (isDark: boolean) =>
    isDark ? 'rounded-full bg-gray-700 p-2' : 'rounded-full bg-gray-100 p-2',
  tabIndicator: 'mt-2 h-[3px] w-8 rounded-full bg-transparent',
};
