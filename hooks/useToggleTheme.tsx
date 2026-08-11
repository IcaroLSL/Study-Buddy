import { useCallback, useEffect, useMemo, useState } from 'react';
import { Appearance, ColorSchemeName } from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { useColorScheme } from 'nativewind';

type ThemeMode = 'light' | 'dark';

const STORAGE_KEY = '@bayoo:theme';

const normalizeScheme = (scheme: ColorSchemeName | null | undefined): ThemeMode => scheme === 'light' ? 'light' : 'dark';


export function useToggleTheme() {
    const { colorScheme, setColorScheme } = useColorScheme();
    const [theme, setTheme] = useState<ThemeMode>(normalizeScheme(colorScheme ?? Appearance.getColorScheme()));

    useEffect(() => {
        (async () => {
            try {
                const saved = await AsyncStorage.getItem(STORAGE_KEY);
                if (saved === 'light' || saved === 'dark') {
                    setTheme(saved);
                    setColorScheme(saved);
                    return;
                }

                const deviceTheme = normalizeScheme(Appearance.getColorScheme());
                setTheme(deviceTheme);
                setColorScheme(deviceTheme);
            } catch {
                const fallbackTheme = normalizeScheme(Appearance.getColorScheme());
                setTheme(fallbackTheme);
                setColorScheme(fallbackTheme);
            }
        })();
    }, [setColorScheme]);

    useEffect(() => {
        const nextTheme = normalizeScheme(colorScheme ?? Appearance.getColorScheme());
        setTheme((currentTheme) => (currentTheme === nextTheme ? currentTheme : nextTheme));
    }, [colorScheme]);

    const toggleTheme = useCallback(async () => {
        const next: ThemeMode = theme === 'light' ? 'dark' : 'light';
        setTheme(next);
        setColorScheme(next);
        await AsyncStorage.setItem(STORAGE_KEY, next);
    }, [theme, setColorScheme]);

    const setLightTheme = useCallback(async () => {
        setTheme('light');
        setColorScheme('light');
        await AsyncStorage.setItem(STORAGE_KEY, 'light');
    }, [setColorScheme]);

    const setDarkTheme = useCallback(async () => {
        setTheme('dark');
        setColorScheme('dark');
        await AsyncStorage.setItem(STORAGE_KEY, 'dark');
    }, [setColorScheme]);

    return useMemo(
        () => ({
            theme,
            isDark: theme === 'dark',
            toggleTheme,
            setLightTheme,
            setDarkTheme,
        }),
        [theme, toggleTheme, setLightTheme, setDarkTheme]
    );
}