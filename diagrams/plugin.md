# Plug-in Requirements

this is my first time doing anything this in depth so I've decided to start making requirements for the microkernal architecture. This could also contain some planned plug-ins but for new I'm focused on design.

## Listed requirements for APIs

For context this app will use Gio for our gui. So far we have outlined:

- plug-in dependencies to determine load order and event listening
- some sort of Plug-in declaration for Ui (is within sidebar, main area, etc.)

for our core, we determined we need:

- UI API
- Event API (dependency based?)