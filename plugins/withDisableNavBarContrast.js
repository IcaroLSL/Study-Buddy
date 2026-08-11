const { withAndroidStyles } = require('@expo/config-plugins');

module.exports = function withDisableNavBarContrast(config) {
  return withAndroidStyles(config, (mod) => {
    const appTheme = mod.modResults.resources.style?.find(
      (s) => s.$.name === 'AppTheme'
    );
    if (!appTheme) {
      console.warn('[withDisableNavBarContrast] AppTheme não encontrado em styles.xml');
      return mod;
    }

    appTheme.item = appTheme.item || [];

    const setItem = (name, value) => {
      const existing = appTheme.item.find((i) => i.$?.name === name);
      if (existing) {
        existing._ = value;
      } else {
        appTheme.item.push({ $: { name }, _: value });
      }
    };

    setItem('android:enforceNavigationBarContrast', 'false');

    return mod;
  });
};
