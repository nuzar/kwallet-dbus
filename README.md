A kwallet dbus call demo.

# Build
```build
flatpak-builder --force-clean --install-deps-from=flathub --user --repo=repo --install builddir com.example.kwallet-dbus.yaml
```

# Execution
run with default args(`org.kde.kwalletd6/modules/kwalletd6:org.kde.KWallet.isEnabled`)
```bash
flatpak run com.example.kwallet-dbus
```


run with dbus `org.freedesktop.secrets` service permission
```bash
flatpak run --talk-name=org.freedesktop.secrets com.example.kwallet-dbus -object org.freedesktop.secrets
```
