---
layout: post
title: "[UE4] Getting Multiplayer Player Pawn AI Navigation to work (C++)"
date: 2014-08-25 15:52:34.000000000 +01:00
type: post
parent_id: '0'
published: true
password: ''
status: publish
categories:
- Game Development
tags:
- UE4
meta:
  _rest_api_published: '1'
  _rest_api_client_id: "-1"
  _publicize_job_id: '5186049052'
  _elasticsearch_data_sharing_indexed_on: '2024-11-18 14:55:00'
permalink: "/2014/08/25/ue4-getting-multiplayer-player-pawn-ai-navigation-to-work-c/"
---

Unreal Engine is an awesome piece of technology making it easy to do
almost anything you might want.

When using the Top Down view however, there is a hurdle to get over when
trying to get multiplayer to work. This is a C++ project solution to
this problem based on a [BluePrints
solution](https://answers.unrealengine.com/questions/34074/does-ue4-have-client-side-prediction-built-in.html).

The basic problem stems from the fact that

> \"*SimpleMoveToLocation* was never intended to be used in a network
> environment. It\'s simple after all ;) Currently there\'s no dedicated
> engine way of making player pawn follow a path. \" (from the same
> page)

To be able to get a working version of *SimpleMoveToLocation*, we need
to do the following:

-   Create a proxy player class (BP_WarriorProxy is BP solution)
-   Set the proxy class as the default player controller class
-   Move the camera to the proxy (Since the actual player class is on
    the server, we can\'t put a camera on it to display on the client)

The BP solution talks about four classes - our counterparts are as
follows:

-   BP_WarriorProxy: ADemoPlayerProxy
-   BP_WarriorController: ADemoPlayerController (Auto-created when
    creating a c++ top down project)
-   BP_Warrior: ADemoCharacter (Auto-created when creating a C++ top
    down project)
-   BP_WarriorAI: AAIController - we have no reason to subclass it.

So, from a standard c++ top down project, the only class we need to add
is the ADemoPlayerProxy - so go ahead and do that first.

The first thing we\'ll do is rewire the ADemoGameMode class to
initialise the proxy class instead of the default MyCharacter Blueprint.

 

\[code language=\"cpp\"\]\
ADemoGameMode::ADemoGameMode(const class
FPostConstructInitializeProperties&amp; PCIP) : Super(PCIP)\
{\
// use our custom PlayerController class\
PlayerControllerClass = ADemoPlayerController::StaticClass();

// set default pawn class to our Blueprinted character\
/\* static ConstructorHelpers::FClassFinder\<apawn\>
PlayerPawnBPClass(TEXT(\"/Game/Blueprints/MyCharacter\"));\
if (PlayerPawnBPClass.Class != NULL)\
{\
DefaultPawnClass = PlayerPawnBPClass.Class;\
}\*/

DefaultPawnClass = ADemoPlayerProxy::StaticClass(); }

\[/code\]

 

Our Player Proxy class declaration

\[code language=\"cpp\"\]\
/\* This class will work as a proxy on the client - tracking the
movements of the\
\* real Character on the server side and sending back controls. \*/\
UCLASS() class Demo_API ADemoPlayerProxy : public APawn\
{\
GENERATED_UCLASS_BODY()\
/\*\* Top down camera \*/\
UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera)
TSubobjectPtr\<class ucameracomponent=\"\"\> TopDownCameraComponent;

/\*\* Camera boom positioning the camera above the character \*/\
UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera)
TSubobjectPtr\<class uspringarmcomponent=\"\"\> CameraBoom;

// Needed so we can pick up the class in the constructor and spawn it
elsewhere\
TSubclassOf\<aactor\> CharacterClass;

// Pointer to the actual character. We replicate it so we know its
location for the camera on the client\
UPROPERTY(Replicated) ADemoCharacter\* Character;

// The AI Controller we will use to auto-navigate the player\
AAIController\* PlayerAI;

// We spawn the real player character and other such elements here\
virtual void BeginPlay() override;

// Use do keep this actor in sync with the real one\
void Tick(float DeltaTime);

// Used by the controller to get moving to work\
void MoveToLocation(const ADemoPlayerController\* controller, const
FVector&amp; vector);\
};

\[/code\]

and the definition:

\[code language=\"cpp\"\]

#include \"Demo.h\"\
#include \"DemoCharacter.h\"\
#include \"AIController.h\"\
#include \"DemoPlayerProxy.h\"\
#include \"UnrealNetwork.h\"

ADemoPlayerProxy::ADemoPlayerProxy(const class
FPostConstructInitializeProperties&amp; PCIP)\
: Super(PCIP)\
{\
// Don\'t rotate character to camera direction\
bUseControllerRotationPitch = false;\
bUseControllerRotationYaw = false;\
bUseControllerRotationRoll = false;

// It seems that without a RootComponent, we can\'t place the Actual
Character easily\
TSubobjectPtr&lt;UCapsuleComponent&gt; TouchCapsule =
PCIP.CreateDefaultSubobject\<ucapsulecomponent\>(this,
TEXT(\"dummy\"));\
TouchCapsule-&gt;InitCapsuleSize(1.0f, 1.0f);\
TouchCapsule-&gt;SetCollisionEnabled(ECollisionEnabled::NoCollision);\
TouchCapsule-&gt;SetCollisionResponseToAllChannels(ECR_Ignore);\
RootComponent = TouchCapsule;

// Create a camera boom\...\
CameraBoom =
PCIP.CreateDefaultSubobject&lt;USpringArmComponent&gt;(this,
TEXT(\"CameraBoom\"));\
CameraBoom-&gt;AttachTo(RootComponent);\
CameraBoom-&gt;bAbsoluteRotation = true; // Don\'t want arm to rotate
when character does\
CameraBoom-&gt;TargetArmLength = 800.f;\
CameraBoom-&gt;RelativeRotation = FRotator(-60.f, 0.f, 0.f);\
CameraBoom-&gt;bDoCollisionTest = false; // Don\'t want to pull camera
in when it collides with level

// Create a camera\...\
TopDownCameraComponent =
PCIP.CreateDefaultSubobject&lt;UCameraComponent&gt;(this,
TEXT(\"TopDownCamera\"));\
TopDownCameraComponent-&gt;AttachTo(CameraBoom,
USpringArmComponent::SocketName);\
TopDownCameraComponent-&gt;bUseControllerViewRotation = false; // Camera
does not rotate relative to arm

if (Role == ROLE_Authority)\
{\
static ConstructorHelpers::FObjectFinder&lt;UClass&gt;
PlayerPawnBPClass(TEXT(\"/Game/Blueprints/MyCharacter.MyCharacter_C\"));\
CharacterClass = PlayerPawnBPClass.Object;\
}

}

void ADemoPlayerProxy::BeginPlay()\
{\
Super::BeginPlay();\
if (Role == ROLE_Authority)\
{\
// Get current location of the Player Proxy\
FVector Location = GetActorLocation();\
FRotator Rotation = GetActorRotation();

FActorSpawnParameters SpawnParams;\
SpawnParams.Owner = this;\
SpawnParams.Instigator = Instigator;\
SpawnParams.bNoCollisionFail = true;

// Spawn the actual player character at the same location as the Proxy\
Character =
Cast&lt;ADemoCharacter&gt;(GetWorld()-&gt;SpawnActor(CharacterClass,
&amp;Location, &amp;Rotation, SpawnParams));

// We use the PlayerAI to control the Player Character for Navigation\
PlayerAI =
GetWorld()-&gt;SpawnActor&lt;AAIController&gt;(GetActorLocation(),
GetActorRotation());\
PlayerAI-&gt;Possess(Character);\
}

}

void ADemoPlayerProxy::Tick(float DeltaTime)\
{

Super::Tick(DeltaTime);\
if (Character)\
{\
// Keep the Proxy in sync with the real character\
FTransform CharTransform = Character-&gt;GetTransform();\
FTransform MyTransform = GetTransform();

FTransform Transform;\
Transform.LerpTranslationScale3D(CharTransform, MyTransform,
ScalarRegister(0.5f));

SetActorTransform(Transform);

}\
}

void ADemoPlayerProxy::MoveToLocation(const ADemoPlayerController\*
controller, const FVector&amp; DestLocation)\
{\
/\*\* Looks easy - doesn\'t it.\
\* What this does is to engage the AI to pathfind.\
\* The AI will then \"route\" the character correctly.\
\* The Proxy (and with it the camera), on each tick, moves to the
location of the real character\
\*\
\* And thus, we get the illusion of moving with the Player Character\
\*/\
PlayerAI-&gt;MoveToLocation(DestLocation);\
}

void ADemoPlayerProxy::GetLifetimeReplicatedProps(TArray&lt; class
FLifetimeProperty &gt; &amp; OutLifetimeProps) const\
{

Super::GetLifetimeReplicatedProps(OutLifetimeProps);

// Replicate to Everyone\
DOREPLIFETIME(ADemoPlayerProxy, Character);\
}\</ucapsulecomponent\>\
\[/code\]

We\'ll now cover changes to the Player Controller. We\'ll rewire it here
to ask the proxy to move, which will in turn ask the AIController to
find a path and move the real player character. 

This involves changing the *SetMoveDestination* method to call a server
side method to move the character. When the character moves, the player
Proxy will automatically mirror the movements.

In the header file, add the following function

\[code language=\"cpp\"\]

/\*\* Navigate player to the given world location (Server Version) \*/\
UFUNCTION(reliable, server, WithValidation)\
void ServerSetNewMoveDestination(const FVector DestLocation);

\[/code\]

Now let\'s implement it (DemoPlayerController.cpp):

\[code language=\"cpp\"\]

bool ADemoPlayerController::ServerSetNewMoveDestination_Validate(const
FVector DestLocation)\
{\
return true;\
}

/\* Actual implementation of the ServerSetMoveDestination method \*/\
void
ADemoPlayerController::ServerSetNewMoveDestination_Implementation(const
FVector DestLocation)\
{\
ADemoPlayerProxy\* Pawn = Cast\<ademoplayerproxy\>(GetPawn());\
if (Pawn)\
{\
UNavigationSystem\* const NaDemoys =
GetWorld()-&gt;GetNavigationSystem();\
float const Distance = FVector::Dist(DestLocation,
Pawn-&gt;GetActorLocation());

// We need to issue move command only if far enough in order for walk
animation to play correctly\
if (NaDemoys &amp;&amp; (Distance &gt; 120.0f))\
{\
//NaDemoys-&gt;SimpleMoveToLocation(this, DestLocation);\
Pawn-&gt;MoveToLocation(this, DestLocation);\
}\
}

}

\[/code\]

And finally, the rewiring of the original method:

\[code language=\"cpp\"\]\
void ADemoPlayerController::SetNewMoveDestination(const FVector
DestLocation)\
{\
ServerSetNewMoveDestination(DestLocation);\
}\
\[/code\]

 

Finally, in terms of the character class, the only change is really to
remove the camera components that we moved to the Player Proxy which I
shall leave to you :-p
