---
categories:
- Game Development
date: "2014-08-25T15:52:34Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:00"
  _publicize_job_id: "5186049052"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
parent_id: "0"
password: ""
published: true
status: publish
tags:
- UE4
title: '[UE4] Getting Multiplayer Player Pawn AI Navigation to work (C++)'
type: post
url: /2014/08/25/ue4-getting-multiplayer-player-pawn-ai-navigation-to-work-c/
---

Unreal Engine is an awesome piece of technology making it easy to do almost
anything you might want.

When using the Top Down view however, there is a hurdle to get over when trying
to get multiplayer to work. This is a C++ project solution to this problem based
on
a [BluePrints solution](https://answers.unrealengine.com/questions/34074/does-ue4-have-client-side-prediction-built-in.html).

The basic problem stems from the fact that

> "_SimpleMoveToLocation_ was never intended to be used in a network
> environment. It's simple after all ;) Currently there's no dedicated engine
> way of making player pawn follow a path. " (from the same page)

To be able to get a working version of _SimpleMoveToLocation_, we need to do the
following:

<!--more-->

- Create a proxy player class (BP_WarriorProxy is BP solution)
- Set the proxy class as the default player controller class
- Move the camera to the proxy (Since the actual player class is on the server,
  we can't put a camera on it to display on the client)

The BP solution talks about four classes - our counterparts are as follows:

- BP_WarriorProxy: ADemoPlayerProxy
- BP_WarriorController: ADemoPlayerController (Auto-created when creating a c++
  top down project)
- BP_Warrior: ADemoCharacter (Auto-created when creating a C++ top down project)
- BP_WarriorAI: AAIController - we have no reason to subclass it.

So, from a standard c++ top down project, the only class we need to add is the
ADemoPlayerProxy - so go ahead and do that first.

The first thing we'll do is rewire the ADemoGameMode class to initialise the
proxy class instead of the default MyCharacter Blueprint.

```c
ADemoGameMode::ADemoGameMode(const class FPostConstructInitializeProperties&
PCIP) : Super(PCIP)
    {
    // use our custom PlayerController class
    PlayerControllerClass = ADemoPlayerController::StaticClass();

    // set default pawn class to our Blueprinted character /_ static
    ConstructorHelpers::FClassFinder<apawn>
    PlayerPawnBPClass(TEXT("/Game/Blueprints/MyCharacter"));
    if (PlayerPawnBPClass.Class != NULL)
    {
        DefaultPawnClass = PlayerPawnBPClass.Class;
    }

    DefaultPawnClass = ADemoPlayerProxy::StaticClass();
}
```

Our Player Proxy class declaration

```cpp
/* This class will work as a proxy on the client - tracking
   the movements of the real Character on the server side
   and sending back controls. */
UCLASS() class Demo_API ADemoPlayerProxy : public APawn
{
   GENERATED_UCLASS_BODY()
        /** Top down camera */
    UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera) TSubobjectPtr<UCameraComponent> TopDownCameraComponent;

        /* Camera boom positioning the camera above the character */
    UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera)
    TSubobjectPtr<USpringArmComponent=""> CameraBoom;

    // Needed so we can pick up the class in the constructor and spawn it elsewhere
    TSubclassOf<AActor> CharacterClass;

    // Pointer to the actual character. We replicate it so we know its
    //location for the camera on the client
    UPROPERTY(Replicated) ADemoCharacter* Character;

    // The AI Controller we will use to auto-navigate the player AAIController\*
    PlayerAI;

    // We spawn the real player character and other such elements here virtual void
    BeginPlay() override;

    // Use do keep this actor in sync with the real one void Tick(float DeltaTime);

    // Used by the controller to get moving to work void MoveToLocation(const
    ADemoPlayerController* controller, const FVector& vector);
};
```

and the definition:

```cpp
#include "Demo.h"
#include "DemoCharacter.h"
#include "AIController.h"
#include "DemoPlayerProxy.h"
#include "UnrealNetwork.h"

ADemoPlayerProxy::ADemoPlayerProxy(const class FPostConstructInitializeProperties& PCIP) : Super(PCIP)
{
    // Don't rotate character to camera direction
    bUseControllerRotationPitch = false;
    bUseControllerRotationYaw = false;
    bUseControllerRotationRoll = false;

    // It seems that without a RootComponent, we can't place the Actual Character easily
    TSubobjectPtr<UCapsuleComponent> TouchCapsule = PCIP.CreateDefaultSubobject<ucapsulecomponent>(this, TEXT("dummy"));
    TouchCapsule->InitCapsuleSize(1.0f, 1.0f);
    TouchCapsule->SetCollisionEnabled(ECollisionEnabled::NoCollision);
    TouchCapsule->SetCollisionResponseToAllChannels(ECR_Ignore);
    RootComponent = TouchCapsule;

    // Create a camera boom... CameraBoom =
    PCIP.CreateDefaultSubobject<USpringArmComponent>(this, TEXT("CameraBoom"));
    CameraBoom->AttachTo(RootComponent);
    CameraBoom->bAbsoluteRotation = true; // Don't want arm to rotate when
    character does CameraBoom->TargetArmLength = 800.f;
    CameraBoom->RelativeRotation = FRotator(-60.f, 0.f, 0.f);
    CameraBoom->bDoCollisionTest = false; // Don't want to pull camera in when it collides with level

    // Create a camera...
    TopDownCameraComponent = PCIP.CreateDefaultSubobject<UCameraComponent>(this, TEXT("TopDownCamera"));
    TopDownCameraComponent->AttachTo(CameraBoom, USpringArmComponent::SocketName);
    TopDownCameraComponent->bUseControllerViewRotation = false; // Camera does not rotate relative to arm

    if (Role == ROLE_Authority)
    {
        static ConstructorHelpers::FObjectFinder<UClass> PlayerPawnBPClass(TEXT("/Game/Blueprints/MyCharacter.MyCharacter_C"));
        CharacterClass = PlayerPawnBPClass.Object;
    }

}

void ADemoPlayerProxy::BeginPlay()
{
    Super::BeginPlay();
    if (Role == ROLE_Authority)
    {
        // Get current location of the Player Proxy
        FVector Location = GetActorLocation(); FRotator Rotation = GetActorRotation();

        FActorSpawnParameters SpawnParams;
        SpawnParams.Owner = this;
        SpawnParams.Instigator = Instigator;
        SpawnParams.bNoCollisionFail = true;

        // Spawn the actual player character at the same location as the Proxy
        Character = Cast<ADemoCharacter>(GetWorld()->SpawnActor(CharacterClass, &Location, &Rotation, SpawnParams));

        // We use the PlayerAI to control the Player Character for Navigation
        PlayerAI = GetWorld()->SpawnActor<AAIController>(GetActorLocation(), GetActorRotation()); PlayerAI->Possess(Character);
    }

}

void ADemoPlayerProxy::Tick(float DeltaTime) {

    Super::Tick(DeltaTime);
    if (Character)
    {
        // Keep the Proxy in sync with the real character
        FTransform CharTransform = Character->GetTransform();
        FTransform MyTransform = GetTransform();

        FTransform Transform;
        Transform.LerpTranslationScale3D(CharTransform, MyTransform, ScalarRegister(0.5f));

        SetActorTransform(Transform);

} }

void ADemoPlayerProxy::MoveToLocation(const ADemoPlayerController* controller, const FVector& DestLocation)
{
    /** Looks easy - doesn't it.

    - What this does is to engage the AI to pathfind.
    - The AI will then "route" the character correctly.
    - The Proxy (and with it the camera), on each tick, moves to the location of the
      real character
    -
    - And thus, we get the illusion of moving with the Player Character */

      PlayerAI->MoveToLocation(DestLocation);
}

void ADemoPlayerProxy::GetLifetimeReplicatedProps(TArray< class FLifetimeProperty > & OutLifetimeProps) const
{

    Super::GetLifetimeReplicatedProps(OutLifetimeProps);

    // Replicate to Everyone
    DOREPLIFETIME(ADemoPlayerProxy, Character);
}
```

We'll now cover changes to the Player Controller. We'll rewire it here to ask
the proxy to move, which will in turn ask the AIController to find a path and
move the real player character.

This involves changing the `SetMoveDestination` method to call a server side
method to move the character. When the character moves, the player Proxy will
automatically mirror the movements.

In the header file, add the following function

```cpp
/*_ Navigate player to the given world location (Server Version) */
UFUNCTION(reliable, server, WithValidation) void
ServerSetNewMoveDestination(const FVector DestLocation);
```

Now let's implement it (DemoPlayerController.cpp):

```
bool ADemoPlayerController::ServerSetNewMoveDestination_Validate(const FVector
DestLocation)
{
    return true;
}

/* Actual implementation of the ServerSetMoveDestination method */
void ADemoPlayerController::ServerSetNewMoveDestination_Implementation(const
FVector DestLocation)
{
    ADemoPlayerProxy* Pawn = Cast<ademoplayerproxy>(GetPawn());
    if (Pawn)
    {
        UNavigationSystem* const NaDemoys = GetWorld()->GetNavigationSystem();
        float const Distance = FVector::Dist(DestLocation, Pawn->GetActorLocation());

        // We need to issue move command only if far enough in order for walk animation
        to play correctly
        if (NaDemoys && (Distance > 120.0f))
        {
            //NaDemoys->SimpleMoveToLocation(this, DestLocation);
            Pawn->MoveToLocation(this, DestLocation);
        }
    }

}
```

And finally, the rewiring of the original method:

```cpp
void ADemoPlayerController::SetNewMoveDestination(const FVector DestLocation)
{
    ServerSetNewMoveDestination(DestLocation);
}
```

Finally, in terms of the character class, the only change is really to remove
the camera components that we moved to the Player Proxy which I shall leave to
you :-p
