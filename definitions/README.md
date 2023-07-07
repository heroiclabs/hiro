configuration
=============

This folder contains example configuration files for Hiro's various systems and features.

## Usage

Each Hiro system accepts its own configuration file, which is passed to it at initialisation time.

For example to initialise the Base system:

```
baseSystem := hiro.WithBaseSystem("Hiro-Base.json", true)
_ = hiro.Init(ctx, logger, nk, initializer, baseSystem)
```

## Satori Usage

When preparing overrides for these configurations in Satori, which are then expected to be read by a configured Hiro Personalizer, the same configuration format is expected.

For example to configure an override for the Energy system:

1. Define a Satori Feature Flag named after the system it configures: `hiro-energy`
2. Use the correct configured schema validator: `Hiro-Energy`
3. Set up the configuration override as the flag value, or an appropriate variant.
